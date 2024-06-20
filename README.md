# protoc-gen-ext
Extensions to protobuf generation for Solo.io.

## Usage

To build generated protos for testing, run ```make install-tools && make generated-code -B```.

## Testing

Run `make run-tests`.

## Troubleshooting problems

If you're attempting to regenerate test protos unsuccessfully and see an error like this:
```
[error] failed post-processing: 298:30: missing ',' in argument list (and 10 more errors)
--ext_out: protoc-gen-ext: Plugin failed with status code 1.
make: *** [Makefile:51: generated-code] Error 1
```

Make sure you are running `make generated-code -B` and including the `-B` flag.
