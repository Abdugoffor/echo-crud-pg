package chromium

import (
	"bytes"
	"context"
	"io"
	"log"
	"sync"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

type ChromiumService interface {
	GenSvg(ctx context.Context, url string) (string, error)
	Run(writer io.Writer, url string) (string, error)
	Close()
	Ctx() context.Context
}

type chromiumService struct {
	allocCtx    context.Context
	allocCancel context.CancelFunc

	chromedpCtx    context.Context
	chromedpCancel context.CancelFunc
}

// TODO: I'll remove it after I try my work
func (c *chromiumService) Ctx() context.Context {
	return c.chromedpCtx
}

func New() (ChromiumService, error) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoSandbox,
		chromedp.DisableGPU,
		chromedp.Headless,
	)

	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), opts...)

	chromedpCtx, chromedpCancel := chromedp.NewContext(allocCtx) // chromedp.WithDebugf(log.Printf),

	if err := chromedp.Run(chromedpCtx); err != nil {
		return nil, err
	}

	return &chromiumService{
		allocCtx:       allocCtx,
		allocCancel:    allocCancel,
		chromedpCtx:    chromedpCtx,
		chromedpCancel: chromedpCancel,
	}, nil
}

func (c *chromiumService) Close() {
	c.allocCancel()
	c.chromedpCancel()
}

func (c *chromiumService) GenSvg(_ context.Context, url string) (string, error) {
	var outer string
	{
		now := time.Now()
		ctx, cancel := chromedp.NewContext(c.chromedpCtx)
		defer cancel()

		err := chromedp.Run(ctx,
			chromedp.Navigate("about:blank"),

			chromedp.ActionFunc(func(ctx context.Context) error {
				lctx, cancel := context.WithCancel(ctx)
				defer cancel()
				var wg sync.WaitGroup
				wg.Add(1)
				chromedp.ListenTarget(lctx, func(ev interface{}) {
					if _, ok := ev.(*page.EventLoadEventFired); ok {
						cancel()
						wg.Done()
					}
				})

				frameTree, err := page.GetFrameTree().Do(ctx)
				if err != nil {
					return err
				}

				if err := page.SetDocumentContent(frameTree.Frame.ID, url).Do(ctx); err != nil {
					return err
				}
				wg.Wait()
				return nil
			}),

			chromedp.Click(`#canvas`, chromedp.ByQuery),
			chromedp.WaitVisible(`#svg`),
			chromedp.OuterHTML("#svg", &outer, chromedp.ByQuery),
		)

		if err != nil {
			return "", err
		}

		log.Println(time.Since(now))
	}

	return outer, nil
}

func (c *chromiumService) Run(writer io.Writer, url string) (string, error) {
	now := time.Now()

	// Создаем контекст для chromedp
	ctx, cancel := chromedp.NewContext(c.chromedpCtx)
	defer cancel()

	// Основная задача
	if err := chromedp.Run(ctx,
		chromedp.Navigate("about:blank"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			lctx, cancel := context.WithCancel(ctx)
			defer cancel()
			var once sync.Once

			chromedp.ListenTarget(lctx, func(ev any) {
				if _, ok := ev.(*page.EventLoadEventFired); ok {
					once.Do(func() {
						cancel()
					})
				}
			})

			frameTree, err := page.GetFrameTree().Do(ctx)
			{
				if err != nil {
					return err
				}

				if err := page.SetDocumentContent(frameTree.Frame.ID, url).Do(ctx); err != nil {
					return err
				}
			}

			return nil
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().
				WithPrintBackground(false).
				WithLandscape(true).
				Do(ctx)
			if err != nil {
				return err
			}

			if _, err := io.Copy(writer, bytes.NewReader(buf)); err != nil {
				return err
			}
			return nil
		}),
	); err != nil {
		return "", err
	}

	// Логируем время выполнения
	log.Printf("Time taken: %v\n", time.Since(now))
	return "", nil
}
