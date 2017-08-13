$tot = 0;

%done = ();

$max = 1000;

$debug = 0;

for $i (0..$max) {
    my $i3 = $i * 3;
    my $i5 = $i * 5;
    if ($i5 < $max) {
        $tot += $i5;
        $done{$i5} = 1;
        print "[5] got " . ($i5) . "\n" if $debug;
    }
    if ($i3 < $max && !$done{$i3}) {
        $tot += $i3;
        $done{$i3} = 1;
        print "[3] got " . ($i3) . "\n" if $debug;
    }
}

print "answer: $tot\n";
