# libAvroPhonetic

Work in progress.

The goal is to create a completely rewritten reference implementation of classic (static) and dictionary based transliterators for Avro Phonetic.

Design goals:
- Cross platform, reusable lib.
- Performant. Avoid regular-expression based search on runtime for generating dictionary based suggestions.
- Good documentations and test coverage. Easier to understand and contribute.
