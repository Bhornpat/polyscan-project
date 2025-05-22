#!/usr/bin/perl
use strict;
use warnings;
use JSON;
# Perl script for regex scanning
my $file = shift;
open(FH, '<', $file) or die "Cannot open file: $!";
my @lines = <FH>;
close(FH);

my @matches = grep /password|token/i, @lines;
print encode_json(\@matches);

