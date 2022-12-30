# KCache

KCache is a set membership object designed for fast look-ups, while maintaining a relatively small memory footprint.

## Background

It was created under the following constraints:

- False Positives (FP) were intolerable, so a Bloom Filter (and derivative systems) were ruled out;
- The system would be populated with fixed-length, uniformly distributed IDs;
- The system would be read-intensive, but rarely written; and
- Ideally, sub-microsecond look-up is desirable.

The built-in `map[string]struct{}` object incurred significant overhead when loading items. On a test of 120M 11-byte ID strings (a 1.3 Gb file), the `map` set was about ~3 Gb in memory. Consequently, a custom solution was pursued that would slim this down while maintaining near-constant time look-up.

## Implementation

This implementation splits up keys into a "prefix" and a "value". The "prefix" serves as a bucket for values, which maps to a simple, sorted slice of byte arrays. Membership is tested via binary search.

On the 120M ID file, the memory footprint is around 1300 Mb (equivalent to the original file), with around a benchmarked random look-up speed of 48-50ns. This equates to roughly ~20,000,000 look-ups/second.

Overall, the performance of the system is exceptionally good, but the original design constraints make KCache a somewhat bespoke solution that should be used carefully.

