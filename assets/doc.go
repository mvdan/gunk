//go:generate protoc -Ibundled/ --include_imports -ogenerated/google_api_annotations.fdp bundled/google/api/annotations.proto
//go:generate protoc -Ibundled/ --include_imports -ogenerated/google_protobuf_empty.fdp bundled/google/protobuf/empty.proto
//go:generate protoc -Ibundled/ --include_imports -ogenerated/google_protobuf_timestamp.fdp bundled/google/protobuf/timestamp.proto
//go:generate protoc -Ibundled/ --include_imports -ogenerated/google_protobuf_duration.fdp bundled/google/protobuf/duration.proto
//go:generate vfsgendev -source="github.com/gunk/gunk/assets".Assets
package assets
