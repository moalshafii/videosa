package video

import "fmt"

func (opts *Options) SubtitleFilter() string {
	// Build subtitle filter with customization options
	return fmt.Sprintf(
		"subtitles=%s:force_style='FontName=%s,FontSize=%d,PrimaryColour=%s,BackColour=%s,Outline=%d,Alignment=%d'",
		opts.SubsFile,
		opts.SubFontName,
		opts.SubFontSize,
		opts.SubFontColor,
		opts.SubBGColor,
		opts.SubBorderWidth,
		opts.SubAlignment,
	)
}

func (opts *Options) OverlayFilter() string {
	return fmt.Sprintf("overlay=%d:%d", opts.Overlay.X, opts.Overlay.Y)
}
