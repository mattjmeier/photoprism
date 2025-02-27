package media

import "github.com/photoprism/photoprism/pkg/fs"

// Formats maps file formats to general media types.
var Formats = map[fs.Type]Type{
	fs.ImageRaw:        Raw,
	fs.ImageDNG:        Raw,
	fs.ImageJPEG:       Image,
	fs.ImageJPEGXL:     Image,
	fs.ImagePNG:        Image,
	fs.ImageGIF:        Image,
	fs.ImageTIFF:       Image,
	fs.ImagePSD:        Image,
	fs.ImageBMP:        Image,
	fs.ImageMPO:        Image,
	fs.ImageAVIF:       Image,
	fs.ImageAVIFS:      Image,
	fs.ImageHEIC:       Image,
	fs.ImageHEICS:      Image,
	fs.VideoHEVC:       Video,
	fs.ImageWebP:       Image,
	fs.VideoWebM:       Video,
	fs.VideoAVI:        Video,
	fs.VideoAVC:        Video,
	fs.VideoVVC:        Video,
	fs.VideoAV1:        Video,
	fs.VideoMPG:        Video,
	fs.VideoMJPG:       Video,
	fs.VideoMP2:        Video,
	fs.VideoMP4:        Video,
	fs.VideoMKV:        Video,
	fs.VideoMOV:        Video,
	fs.Video3GP:        Video,
	fs.Video3G2:        Video,
	fs.VideoFlash:      Video,
	fs.VideoAVCHD:      Video,
	fs.VideoBDAV:       Video,
	fs.VideoOGV:        Video,
	fs.VideoASF:        Video,
	fs.VideoWMV:        Video,
	fs.VectorSVG:       Vector,
	fs.VectorAI:        Vector,
	fs.VectorPS:        Vector,
	fs.VectorEPS:       Vector,
	fs.SidecarXMP:      Sidecar,
	fs.SidecarXML:      Sidecar,
	fs.SidecarAAE:      Sidecar,
	fs.SidecarYAML:     Sidecar,
	fs.SidecarText:     Sidecar,
	fs.SidecarJSON:     Sidecar,
	fs.SidecarMarkdown: Sidecar,
	fs.UnknownType:     Other,
}
