all: carton-install lib/MyService.pm

carton-install:
	cpanm -L local --notest IO::Socket::SSL && carton install

lib/MyService.pm: ../service.proto
	mkdir -p $(dir $@)
	carton exec -- protoc --perl-gpd_out=package=MyService:$(dir $@) -I$(dir $<) $(notdir $<)

clean:
	rm -rf lib/MyService.pm

test:
	carton exec -- prove -Ilib -r t

.PHONY: all carton-install clean test
