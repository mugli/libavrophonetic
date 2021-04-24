package rulebasedcnv

// To make things simpler, Avro Phonetic tries to produce output always in the NFC form (See Unicode normalization).
// Code Editors often break it to canonically equivalent NFD form. To avoid that decomposing, we define
// these 3 Bengali characters like this. Rest of the Bengali characters don't have to deal with it.
const bnRRA = "\u09dc" // ড়
const bnRHA = "\u09dd" // ঢ়
const bnYYA = "\u09df" // য়
