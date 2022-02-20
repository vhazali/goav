// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"unsafe"

	"github.com/vhazali/goav/avcodec"
	"github.com/vhazali/goav/avutil"
)

func (avs *Stream) CodecParameters() *avcodec.AvCodecParameters {
	return (*avcodec.AvCodecParameters)(unsafe.Pointer(avs.codecpar))
}

func (avs *Stream) Codec() *CodecContext {
	return (*CodecContext)(unsafe.Pointer(avs.Codec()))
}

func (avs *Stream) Metadata() *avutil.Dictionary {
	return (*avutil.Dictionary)(unsafe.Pointer(avs.metadata))
}

// Deprecated: move AVStream.*index_entries* to AVStreamInternal
// Commit: cea7c19cda0ea1630ae1de8c102ab14231b9db10
// func (avs *Stream) IndexEntries() *AvIndexEntry {
// 	return (*AvIndexEntry)(unsafe.Pointer(avs.index_entries))
// }

func (avs *Stream) AttachedPic() avcodec.Packet {
	return *fromCPacket(&avs.attached_pic)
}

func (avs *Stream) SideData() *AvPacketSideData {
	return (*AvPacketSideData)(unsafe.Pointer(avs.side_data))
}

// Deprecated
// could not find commit
// func (avs *Stream) ProbeData() AvProbeData {
// 	return AvProbeData(avs.probe_data)
// }

func (avs *Stream) AvgFrameRate() avcodec.Rational {
	return newRational(avs.avg_frame_rate)
}

// func (avs *Stream) DisplayAspectRatio() *Rational {
// 	return (*Rational)(unsafe.Pointer(avs.display_aspect_ratio))
// }

func (avs *Stream) RFrameRate() avcodec.Rational {
	return newRational(avs.r_frame_rate)
}

func (avs *Stream) SampleAspectRatio() avcodec.Rational {
	return newRational(avs.sample_aspect_ratio)
}

func (avs *Stream) TimeBase() avcodec.Rational {
	return newRational(avs.time_base)
}

// func (avs *Stream) RecommendedEncoderConfiguration() string {
// 	return C.GoString(avs.recommended_encoder_configuration)
// }

func (avs *Stream) Discard() AvDiscard {
	return AvDiscard(avs.discard)
}

// Deprecated
// could not find commit
// func (avs *Stream) NeedParsing() AvStreamParseType {
// 	return AvStreamParseType(avs.need_parsing)
// }

// Deprecated: move AVStream.codec_info_nb_frames to AVStreamInternal
// Commit: 7489f632815c98ad58c3db71d1a5239b5dae266c
// func (avs *Stream) CodecInfoNbFrames() int {
// 	return int(avs.codec_info_nb_frames)
// }

func (avs *Stream) Disposition() int {
	return int(avs.disposition)
}

func (avs *Stream) EventFlags() int {
	return int(avs.event_flags)
}

func (avs *Stream) Id() int {
	return int(avs.id)
}

func (avs *Stream) Index() int {
	return int(avs.index)
}

// Deprecated: move AVStream.{inject_global_side_data,display_aspect_ratio}
// to AVStreamInternal
// Commit: c1b916580ae92abca583d9afa2f9f64165292dd8
// func (avs *Stream) InjectGlobalSideData() int {
// 	return int(avs.inject_global_side_data)
// }

// Deprecated: move AVStream.last-IP_{pts,duration} to AVStreamInternal
// Commit: 1262f67cca1ed045521dc3af0d9d0ee9716335f2
// func (avs *Stream) LastIpDuration() int {
// 	return int(avs.last_IP_duration)
// }

// Deprecated: move AVStream.{nb_decoded_frames,mux_ts_offset}
// to AVStreamInternal
// Commit: cb46a6bcbcb85b3910cc0cce78399686ab8efff6
// func (avs *Stream) NbDecodedFrames() int {
// 	return int(avs.nb_decoded_frames)
// }

// Deprecated
// could not find commit
// func (avs *Stream) NbIndexEntries() int {
// 	return int(avs.nb_index_entries)
// }

func (avs *Stream) NbSideData() int {
	return int(avs.nb_side_data)
}

// Deprecated: avformat: move AVStream.probe_packets to AVStreamInternal
// Commit: fab2ed47042d4cc2a4cd69bb97738024c01300c7
// func (avs *Stream) ProbePackets() int {
// 	return int(avs.probe_packets)
// }

// Deprecated
// could not find commit
// func (avs *Stream) PtsWrapBehavior() int {
// 	return int(avs.pts_wrap_behavior)
// }

// Deprecated: move AVStream.{request_probe,skip_to_keyframe} to AVStreamInternal
// Commit: 108864acee1d7b0cde653cee113f3001f1b8915a
// func (avs *Stream) RequestProbe() int {
// 	return int(avs.request_probe)
// }

// Deprecated: move AVStream.{*skip_samples.*_discard_sample} to AVStreamInternal
// Commit: 456b170bd747ea7181c7305fd45278ea251f45ab
// func (avs *Stream) SkipSamples() int {
// 	return int(avs.skip_samples)
// }

// Deprecated: move AVStream.{request_probe,skip_to_keyframe} to AVStreamInternal
// Commit: 108864acee1d7b0cde653cee113f3001f1b8915a
// func (avs *Stream) SkipToKeyframe() int {
// 	return int(avs.skip_to_keyframe)
// }

// Deprecated: move AVStream.stream_identifier to AVStreamInternal
// Commit: f140239777270161bfdf3964bfb057b43ad66a1a
// func (avs *Stream) StreamIdentifier() int {
// 	return int(avs.stream_identifier)
// }

// Deprecated: move AVStream.{pts_wrap_*,update_initial_durations_done}
// to AVStreamInternal
// Commit: 323c9a8c523d772d7439cf386749f0243fecfa9a
// func (avs *Stream) UpdateInitialDurationsDone() int {
// 	return int(avs.update_initial_durations_done)
// }

// Deprecated
// could not find commit
// func (avs *Stream) CurDts() int64 {
// 	return int64(avs.cur_dts)
// }

func (avs *Stream) Duration() int64 {
	return int64(avs.duration)
}

// func (avs *Stream) FirstDiscardSample() int64 {
// 	return int64(avs.first_discard_sample)
// }

// Deprecated: move AVStream.{first,cur}_dts to AVStreamInternal
// Commit: 591b88e6787c4e678237f02a50421d101abd25c2
// func (avs *Stream) FirstDts() int64 {
// 	return int64(avs.first_dts)
// }

// Deprecated
// could not find commit
// func (avs *Stream) InterleaverChunkDuration() int64 {
// 	return int64(avs.interleaver_chunk_duration)
// }

// Deprecated
// could not find commit
// func (avs *Stream) InterleaverChunkSize() int64 {
// 	return int64(avs.interleaver_chunk_size)
// }

// func (avs *Stream) LastDiscardSample() int64 {
// 	return int64(avs.last_discard_sample)
// }

// Deprecated: move AVStream.{last_dts_for_order_check,dts_[mis]ordered}
// to AVStreamInternal
// Commit: 25bade3258af48ba309cee5a0c668a8c70334d04
// func (avs *Stream) LastDtsForOrderCheck() int64 {
// 	return int64(avs.last_dts_for_order_check)
// }

// Deprecated: move AVStream.last-IP_{pts,duration} to AVStreamInternal
// Commit: 1262f67cca1ed045521dc3af0d9d0ee9716335f2
// func (avs *Stream) LastIpPts() int64 {
// 	return int64(avs.last_IP_pts)
// }

// Deprecated: move AVStream.{nb_decoded_frames,mux_ts_offset}
// to AVStreamInternal
// Commit: cb46a6bcbcb85b3910cc0cce78399686ab8efff6
// func (avs *Stream) MuxTsOffset() int64 {
// 	return int64(avs.mux_ts_offset)
// }

func (avs *Stream) NbFrames() int64 {
	return int64(avs.nb_frames)
}

// Deprecated: move AVStream.pts_buffer to AVStreamInternal
// Commit: 91a98cc4ead6d871053011dafb72458a4e75fa8b
// func (avs *Stream) PtsBuffer() int64 {
// 	return int64(avs.pts_buffer[0])
// }

// Deprecated: move AVStream.pts_reorder_error[_count] to AVStreamInternal
// Commit: 36d7c1dee8d1fa1d81130932d2df93a8866c535f
// func (avs *Stream) PtsReorderError() int64 {
// 	return int64(avs.pts_reorder_error[0])
// }

// Deprecated
// could not find commit
// func (avs *Stream) PtsWrapReference() int64 {
// 	return int64(avs.pts_wrap_reference)
// }

// func (avs *Stream) StartSkipSamples() int64 {
// 	return int64(avs.start_skip_samples)
// }

func (avs *Stream) StartTime() int64 {
	return int64(avs.start_time)
}

func (avs *Stream) Parser() *CodecParserContext {
	return (*CodecParserContext)(unsafe.Pointer(avs.Parser()))
}

// Deprecated: move AVStream.last_in_packet_buffer to AVStreamInternal
// Commit: 744b7f2e91fad0953505c909bc2b0cebced03d28
// func (avs *Stream) LastInPacketBuffer() *AvPacketList {
// 	return (*AvPacketList)(unsafe.Pointer(avs.last_in_packet_buffer))
// }

// func (avs *Stream) PrivPts() *FFFrac {
// 	return (*FFFrac)(unsafe.Pointer(avs.priv_pts))
// }

// Deprecated
// could not find commit
// func (avs *Stream) DtsMisordered() uint8 {
// 	return uint8(avs.dts_misordered)
// }

// Deprecated: move AVStream.{last_dts_for_order_check,dts_[mis]ordered}
// to AVStreamInternal
// Commit: 25bade3258af48ba309cee5a0c668a8c70334d04
// func (avs *Stream) DtsOrdered() uint8 {
// 	return uint8(avs.dts_ordered)
// }

// Deprecated: move AVStream.pts_reorder_error[_count] to AVStreamInternal
// Commit: 36d7c1dee8d1fa1d81130932d2df93a8866c535f
// func (avs *Stream) PtsReorderErrorCount() uint8 {
// 	return uint8(avs.pts_reorder_error_count[0])
// }

// Deprecated
// could not find commit
// func (avs *Stream) IndexEntriesAllocatedSize() uint {
// 	return uint(avs.index_entries_allocated_size)
// }

func (avs *Stream) Free() {
	C.av_freep(unsafe.Pointer(avs))
}
