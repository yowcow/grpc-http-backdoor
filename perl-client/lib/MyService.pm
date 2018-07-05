package MyService;

use strict;
use warnings;
use MIME::Base64 qw();
use Google::ProtocolBuffers::Dynamic;

my $gpd = Google::ProtocolBuffers::Dynamic->new;

$gpd->load_serialized_string(MIME::Base64::decode_base64(<<'EOD'));
Cp0BCg1zZXJ2aWNlLnByb3RvEgdzZXJ2aWNlIkYKBlBlcnNvbhIOCgJpZBgBIAIoBVICaWQSEgoE
bmFtZRgCIAIoCVIEbmFtZRIYCgdhZGRyZXNzGAMgASgJUgdhZGRyZXNzIgYKBFZvaWQyMwoERGF0
YRIrCglHZXRQZXJzb24SDS5zZXJ2aWNlLlZvaWQaDy5zZXJ2aWNlLlBlcnNvbg==

EOD


$gpd->map(
    {
      'package' => 'service',
      'prefix' => 'MyService::Service'
    },
);

undef $gpd;

1;
