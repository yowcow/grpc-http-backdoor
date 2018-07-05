use strict;
use warnings;
use Test::More;

use_ok 'MyService';

subtest 'MyService::Person' => sub {
    my $p = MyService::Service::Person->new;

    isa_ok $p, 'MyService::Service::Person';

    $p->set_id(123);
    $p->set_name('Hoge Fuga');
    $p->set_address('234 Foo Bar');

    my $bytes = MyService::Service::Person->encode($p);
    my $p2 = MyService::Service::Person->decode($bytes);

    is $p2->get_id, 123;
    is $p2->get_name, 'Hoge Fuga';
    is $p2->get_address, '234 Foo Bar';
};

done_testing;
