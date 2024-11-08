//package iterx
//
//import (
//	"iter"
//	"slices"
//)
//
//type SliceSeq[E any] struct {
//	seq iter.Seq2[int, E]
//}
//
//func (s SliceSeq[E]) All() iter.Seq2[int, E] {
//	return s.seq
//}
//
//func (s SliceSeq[E]) Filter(filter func(int, E) bool) SliceSeq[E] {
//	return SliceSeq[E]{
//		seq: func(yield func(int, E) bool) {
//			i := 0
//			for k, v := range s.seq {
//				if filter(k, v) {
//					if !yield(i, v) {
//						return
//					}
//					i++
//				}
//			}
//		},
//	}
//}
//
//func (s SliceSeq[E]) Map(mapFn func(E) E) SliceSeq[E] {
//	return SliceSeq[E]{
//		seq: func(yield func(int, E) bool) {
//			for k, v := range s.seq {
//				if !yield(k, mapFn(v)) {
//					return
//				}
//			}
//		},
//	}
//}
//
//func (s SliceSeq[E]) Fill(fill E) SliceSeq[E] {
//	return SliceSeq[E]{
//		seq: func(yield func(int, E) bool) {
//			for i, _ := range s.seq {
//				if !yield(i, fill) {
//					return
//				}
//			}
//		},
//	}
//}
//
//func (s SliceSeq[E]) Find(equal func(int, E) bool) (_ E) {
//	for i, v := range s.seq {
//		if equal(i, v) {
//			return v
//		}
//	}
//	return
//}
//
//func (s SliceSeq[E]) Some(match func(int, E) bool) bool {
//	for i, v := range s.seq {
//		if match(i, v) {
//			return true
//		}
//	}
//	return false
//}
//
//func (s SliceSeq[E]) Every(match func(int, E) bool) bool {
//	for i, v := range s.seq {
//		if !match(i, v) {
//			return false
//		}
//	}
//	return true
//}
//
//func (s SliceSeq[E]) Collect() []E {
//	var res []E
//	for _, v := range s.seq {
//		res = append(res, v)
//	}
//	return res
//}
//
//func (s SliceSeq[E]) Sort(cmp func(x, y E) int) []E {
//	collect := s.Collect()
//	slices.SortFunc(collect, cmp)
//	return collect
//}
//
//func (s SliceSeq[E]) SortStable(cmp func(x, y E) int) []E {
//	collect := s.Collect()
//	slices.SortStableFunc(collect, cmp)
//	return collect
//}
//
//func Slice[S ~[]E, E any](s S) SliceSeq[E] {
//	return SliceSeq[E]{seq: slices.All(s)}
//}
