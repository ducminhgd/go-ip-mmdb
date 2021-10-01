#!/usr/local/bin/perl
# from https://gist.github.com/codesnik/2e0cd34bd98c7c2e4daf
# run 'mcpan MaxMind::DB::Writer' beforehand

use MaxMind::DB::Writer::Tree;
use Net::Works::Network;

my %types = (
#color => 'utf8_string',
#dogs  => [ 'array', 'utf8_string' ],
#size  => 'uint16',
);

my $tree = MaxMind::DB::Writer::Tree->new(
ip_version            => 6,
record_size           => 28,
database_type         => 'GeoLite2-City',
languages             => ['en', 'ru'],
description           => { en => 'GeoLite2 City database (empty)' },
map_key_type_callback => sub { $types{ $_[0] } },
);

# my $network = Net::Works::Network->new_from_string( string => '2001:db8::/48' );

# $tree->insert_network(
# $network,
# {
#     color => 'blue',
#     dogs  => [ 'Fido', 'Ms. Pretty Paws' ],
#     size  => 42,
# },
# );

open my $fh, '>:raw', 'empty.mmdb';
$tree->write_tree($fh);