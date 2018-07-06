use strict;
use warnings;
use Data::Dumper;
use Hijk;
use MyService;

my $host = 'localhost';
my $port = '9998';
my $path = '/GetPerson/';

run_http_client();

sub run_http_client {
    my $resp = Hijk::request({
        method => 'GET',
        host => $host,
        port => $port,
        path => $path,
    });

    die "Failed connecting to http server: $resp->{error}"
        if exists $resp->{error} and $resp->{error} & Hijk::Error::TIMEOUT;

    die "Got status other than OK: $resp->{status}"
        if $resp->{status} != 200;

    my $p = MyService::Service::Person->decode($resp->{body});
    warn Dumper $p;
}
