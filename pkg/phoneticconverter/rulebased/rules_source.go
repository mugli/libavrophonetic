package rulebased

/*
Note:
Many re-implementation of Avro Classic Phonetic saw the JavaScript implementation
(https://github.com/torifat/jsAvroPhonetic/blob/master/src/avro-lib.js) and thought it might be a better idea to extract
the rules in a separate JSON file to make Avro Phonetic more configurable by the users. But I am going to keep it hardcoded like this.

Reasons:
- In practice, that customization never happened. These rules are complex and may lead to subtle errors and end users
never want that kind of configurability.

- Moreover, these rules are also associated with the muscle memory of the Avro Keyboard (IME) users.
If you arbitrarily change them, they are no longer "Avro Phonetic".

- It's more important to write the test-cases to check if as a whole Avro Phonetic conversion is working as
expected from the user's perspective.
*/

const suffix = "suffix"
const prefix = "prefix"
const vowel = "vowel"
const notVowel = "!vowel"
const consonant = "consonant"
const notConsonant = "!consonant"
const punctuation = "punctuation"
const notPunctuation = "!punctuation"
const exactly = "exactly"
const notExactly = "!exactly"

var avroClassicPhoneticRules = []rule{
	// match-replace rules are most basic form of a rule. If we find this match, we replace it with that in the output
	{
		// match is always case-sensitive. That's why we use fixCase function on the input string before trying to transliterate them.
		match:   "bhl",
		replace: "ভ্ল",
	},
	// More complex rules look like this. When we find a match like this , we first check if any exceptions are true.
	// "thenReplace" from an exception block always gets more priority over plain "replace".
	// If none of the exception conditions are true, it works as basic match-replace rule.
	{
		match:   "a",
		replace: "া",
		exceptions: []exception{
			{
				// ifAllMatch works as "AND" for all the rules inside it. If all of them are true, "thenReplace" if this block takes place
				// If any of the condition is false, we move on to the next exception ifAllMatch block.
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   punctuation,
					},
					// The rule is supposed to read like:
					// (Our current position is "a") and when the suffix (next characters after that) is not exactly "`"
					{
						when:  suffix,
						is:    notExactly,
						value: "`",
					},
				},
				thenReplace: "আ",
			},
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   notConsonant,
					},
					{
						when:  prefix,
						is:    notExactly,
						value: "a",
					},
					{
						when:  suffix,
						is:    notExactly,
						value: "`",
					},
				},
				thenReplace: bnYYA + "া",
			},
			{
				ifAllMatch: []matchCondition{
					{
						when:  prefix,
						is:    exactly,
						value: "a",
					},
					{
						when:  suffix,
						is:    notExactly,
						value: "`",
					},
				},
				thenReplace: "আ",
			},
		},
	},
	{
		match:   "psh",
		replace: "পশ",
	},
	{
		match:   "bdh",
		replace: "ব্ধ",
	},
	{
		match:   "bj",
		replace: "ব্জ",
	},
	{
		match:   "bd",
		replace: "ব্দ",
	},
	{
		match:   "bb",
		replace: "ব্ব",
	},
	{
		match:   "bl",
		replace: "ব্ল",
	},
	{
		match:   "bh",
		replace: "ভ",
	},
	{
		match:   "vl",
		replace: "ভ্ল",
	},
	{
		match:   "b",
		replace: "ব",
	},
	{
		match:   "v",
		replace: "ভ",
	},
	{
		match:   "cNG",
		replace: "চ্ঞ",
	},
	{
		match:   "cch",
		replace: "চ্ছ",
	},
	{
		match:   "cc",
		replace: "চ্চ",
	},
	{
		match:   "ch",
		replace: "ছ",
	},
	{
		match:   "c",
		replace: "চ",
	},
	{
		match:   "dhn",
		replace: "ধ্ন",
	},
	{
		match:   "dhm",
		replace: "ধ্ম",
	},
	{
		match:   "dgh",
		replace: "দ্ঘ",
	},
	{
		match:   "ddh",
		replace: "দ্ধ",
	},
	{
		match:   "dbh",
		replace: "দ্ভ",
	},
	{
		match:   "dv",
		replace: "দ্ভ",
	},
	{
		match:   "dm",
		replace: "দ্ম",
	},
	{
		match:   "DD",
		replace: "ড্ড",
	},
	{
		match:   "Dh",
		replace: "ঢ",
	},
	{
		match:   "dh",
		replace: "ধ",
	},
	{
		match:   "dg",
		replace: "দ্গ",
	},
	{
		match:   "dd",
		replace: "দ্দ",
	},
	{
		match:   "D",
		replace: "ড",
	},
	{
		match:   "d",
		replace: "দ",
	},
	{
		match:   "...",
		replace: "...",
	},
	{
		match:   ".`",
		replace: ".",
	},
	{
		match:   "..",
		replace: "।।",
	},
	{
		match:   ".",
		replace: "।",
	},
	{
		match:   "ghn",
		replace: "ঘ্ন",
	},
	{
		match:   "Ghn",
		replace: "ঘ্ন",
	},
	{
		match:   "gdh",
		replace: "গ্ধ",
	},
	{
		match:   "Gdh",
		replace: "গ্ধ",
	},
	{
		match:   "gN",
		replace: "গ্ণ",
	},
	{
		match:   "GN",
		replace: "গ্ণ",
	},
	{
		match:   "gn",
		replace: "গ্ন",
	},
	{
		match:   "Gn",
		replace: "গ্ন",
	},
	{
		match:   "gm",
		replace: "গ্ম",
	},
	{
		match:   "Gm",
		replace: "গ্ম",
	},
	{
		match:   "gl",
		replace: "গ্ল",
	},
	{
		match:   "Gl",
		replace: "গ্ল",
	},
	{
		match:   "gg",
		replace: "জ্ঞ",
	},
	{
		match:   "GG",
		replace: "জ্ঞ",
	},
	{
		match:   "Gg",
		replace: "জ্ঞ",
	},
	{
		match:   "gG",
		replace: "জ্ঞ",
	},
	{
		match:   "gh",
		replace: "ঘ",
	},
	{
		match:   "Gh",
		replace: "ঘ",
	},
	{
		match:   "g",
		replace: "গ",
	},
	{
		match:   "G",
		replace: "গ",
	},
	{
		match:   "hN",
		replace: "হ্ণ",
	},
	{
		match:   "hn",
		replace: "হ্ন",
	},
	{
		match:   "hm",
		replace: "হ্ম",
	},
	{
		match:   "hl",
		replace: "হ্ল",
	},
	{
		match:   "h",
		replace: "হ",
	},
	{
		match:   "jjh",
		replace: "জ্ঝ",
	},
	{
		match:   "jNG",
		replace: "জ্ঞ",
	},
	{
		match:   "jh",
		replace: "ঝ",
	},
	{
		match:   "jj",
		replace: "জ্জ",
	},
	{
		match:   "j",
		replace: "জ",
	},
	{
		match:   "J",
		replace: "জ",
	},
	{
		match:   "kkhN",
		replace: "ক্ষ্ণ",
	},
	{
		match:   "kShN",
		replace: "ক্ষ্ণ",
	},
	{
		match:   "kkhm",
		replace: "ক্ষ্ম",
	},
	{
		match:   "kShm",
		replace: "ক্ষ্ম",
	},
	{
		match:   "kxN",
		replace: "ক্ষ্ণ",
	},
	{
		match:   "kxm",
		replace: "ক্ষ্ম",
	},
	{
		match:   "kkh",
		replace: "ক্ষ",
	},
	{
		match:   "kSh",
		replace: "ক্ষ",
	},
	{
		match:   "ksh",
		replace: "কশ",
	},
	{
		match:   "kx",
		replace: "ক্ষ",
	},
	{
		match:   "kk",
		replace: "ক্ক",
	},
	{
		match:   "kT",
		replace: "ক্ট",
	},
	{
		match:   "kt",
		replace: "ক্ত",
	},
	{
		match:   "kl",
		replace: "ক্ল",
	},
	{
		match:   "ks",
		replace: "ক্স",
	},
	{
		match:   "kh",
		replace: "খ",
	},
	{
		match:   "k",
		replace: "ক",
	},
	{
		match:   "lbh",
		replace: "ল্ভ",
	},
	{
		match:   "ldh",
		replace: "ল্ধ",
	},
	{
		match:   "lkh",
		replace: "লখ",
	},
	{
		match:   "lgh",
		replace: "লঘ",
	},
	{
		match:   "lph",
		replace: "লফ",
	},
	{
		match:   "lk",
		replace: "ল্ক",
	},
	{
		match:   "lg",
		replace: "ল্গ",
	},
	{
		match:   "lT",
		replace: "ল্ট",
	},
	{
		match:   "lD",
		replace: "ল্ড",
	},
	{
		match:   "lp",
		replace: "ল্প",
	},
	{
		match:   "lv",
		replace: "ল্ভ",
	},
	{
		match:   "lm",
		replace: "ল্ম",
	},
	{
		match:   "ll",
		replace: "ল্ল",
	},
	{
		match:   "lb",
		replace: "ল্ব",
	},
	{
		match:   "l",
		replace: "ল",
	},
	{
		match:   "mth",
		replace: "ম্থ",
	},
	{
		match:   "mph",
		replace: "ম্ফ",
	},
	{
		match:   "mbh",
		replace: "ম্ভ",
	},
	{
		match:   "mpl",
		replace: "মপ্ল",
	},
	{
		match:   "mn",
		replace: "ম্ন",
	},
	{
		match:   "mp",
		replace: "ম্প",
	},
	{
		match:   "mv",
		replace: "ম্ভ",
	},
	{
		match:   "mm",
		replace: "ম্ম",
	},
	{
		match:   "ml",
		replace: "ম্ল",
	},
	{
		match:   "mb",
		replace: "ম্ব",
	},
	{
		match:   "mf",
		replace: "ম্ফ",
	},
	{
		match:   "m",
		replace: "ম",
	},
	{
		match:   "0",
		replace: "০",
	},
	{
		match:   "1",
		replace: "১",
	},
	{
		match:   "2",
		replace: "২",
	},
	{
		match:   "3",
		replace: "৩",
	},
	{
		match:   "4",
		replace: "৪",
	},
	{
		match:   "5",
		replace: "৫",
	},
	{
		match:   "6",
		replace: "৬",
	},
	{
		match:   "7",
		replace: "৭",
	},
	{
		match:   "8",
		replace: "৮",
	},
	{
		match:   "9",
		replace: "৯",
	},
	{
		match:   "NgkSh",
		replace: "ঙ্ক্ষ",
	},
	{
		match:   "Ngkkh",
		replace: "ঙ্ক্ষ",
	},
	{
		match:   "NGch",
		replace: "ঞ্ছ",
	},
	{
		match:   "Nggh",
		replace: "ঙ্ঘ",
	},
	{
		match:   "Ngkh",
		replace: "ঙ্খ",
	},
	{
		match:   "NGjh",
		replace: "ঞ্ঝ",
	},
	{
		match:   "ngOU",
		replace: "ঙ্গৌ",
	},
	{
		match:   "ngOI",
		replace: "ঙ্গৈ",
	},
	{
		match:   "Ngkx",
		replace: "ঙ্ক্ষ",
	},
	{
		match:   "NGc",
		replace: "ঞ্চ",
	},
	{
		match:   "nch",
		replace: "ঞ্ছ",
	},
	{
		match:   "njh",
		replace: "ঞ্ঝ",
	},
	{
		match:   "ngh",
		replace: "ঙ্ঘ",
	},
	{
		match:   "Ngk",
		replace: "ঙ্ক",
	},
	{
		match:   "Ngx",
		replace: "ঙ্ষ",
	},
	{
		match:   "Ngg",
		replace: "ঙ্গ",
	},
	{
		match:   "Ngm",
		replace: "ঙ্ম",
	},
	{
		match:   "NGj",
		replace: "ঞ্জ",
	},
	{
		match:   "ndh",
		replace: "ন্ধ",
	},
	{
		match:   "nTh",
		replace: "ন্ঠ",
	},
	{
		match:   "NTh",
		replace: "ণ্ঠ",
	},
	{
		match:   "nth",
		replace: "ন্থ",
	},
	{
		match:   "nkh",
		replace: "ঙ্খ",
	},
	{
		match:   "ngo",
		replace: "ঙ্গ",
	},
	{
		match:   "nga",
		replace: "ঙ্গা",
	},
	{
		match:   "ngi",
		replace: "ঙ্গি",
	},
	{
		match:   "ngI",
		replace: "ঙ্গী",
	},
	{
		match:   "ngu",
		replace: "ঙ্গু",
	},
	{
		match:   "ngU",
		replace: "ঙ্গূ",
	},
	{
		match:   "nge",
		replace: "ঙ্গে",
	},
	{
		match:   "ngO",
		replace: "ঙ্গো",
	},
	{
		match:   "NDh",
		replace: "ণ্ঢ",
	},
	{
		match:   "nsh",
		replace: "নশ",
	},
	{
		match:   "Ngr",
		replace: "ঙর",
	},
	{
		match:   "NGr",
		replace: "ঞর",
	},
	{
		match:   "ngr",
		replace: "ংর",
	},
	{
		match:   "nj",
		replace: "ঞ্জ",
	},
	{
		match:   "Ng",
		replace: "ঙ",
	},
	{
		match:   "NG",
		replace: "ঞ",
	},
	{
		match:   "nk",
		replace: "ঙ্ক",
	},
	{
		match:   "ng",
		replace: "ং",
	},
	{
		match:   "nn",
		replace: "ন্ন",
	},
	{
		match:   "NN",
		replace: "ণ্ণ",
	},
	{
		match:   "Nn",
		replace: "ণ্ন",
	},
	{
		match:   "nm",
		replace: "ন্ম",
	},
	{
		match:   "Nm",
		replace: "ণ্ম",
	},
	{
		match:   "nd",
		replace: "ন্দ",
	},
	{
		match:   "nT",
		replace: "ন্ট",
	},
	{
		match:   "NT",
		replace: "ণ্ট",
	},
	{
		match:   "nD",
		replace: "ন্ড",
	},
	{
		match:   "ND",
		replace: "ণ্ড",
	},
	{
		match:   "nt",
		replace: "ন্ত",
	},
	{
		match:   "ns",
		replace: "ন্স",
	},
	{
		match:   "nc",
		replace: "ঞ্চ",
	},
	{
		match:   "n",
		replace: "ন",
	},
	{
		match:   "N",
		replace: "ণ",
	},
	{
		match:   "OI`",
		replace: "ৈ",
	},
	{
		match:   "OU`",
		replace: "ৌ",
	},
	{
		match:   "O`",
		replace: "ো",
	},
	{
		match:   "OI",
		replace: "ৈ",
		exceptions: []exception{
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   notConsonant,
					},
				},
				thenReplace: "ঐ",
			},
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   punctuation,
					},
				},
				thenReplace: "ঐ",
			},
		},
	},
	{
		match:   "OU",
		replace: "ৌ",
		exceptions: []exception{
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   notConsonant,
					},
				},
				thenReplace: "ঔ",
			},
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   punctuation,
					},
				},
				thenReplace: "ঔ",
			},
		},
	},
	{
		match:   "O",
		replace: "ো",
		exceptions: []exception{
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   notConsonant,
					},
				},
				thenReplace: "ও",
			},
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   punctuation,
					},
				},
				thenReplace: "ও",
			},
		},
	},
	{
		match:   "phl",
		replace: "ফ্ল",
	},
	{
		match:   "pT",
		replace: "প্ট",
	},
	{
		match:   "pt",
		replace: "প্ত",
	},
	{
		match:   "pn",
		replace: "প্ন",
	},
	{
		match:   "pp",
		replace: "প্প",
	},
	{
		match:   "pl",
		replace: "প্ল",
	},
	{
		match:   "ps",
		replace: "প্স",
	},
	{
		match:   "ph",
		replace: "ফ",
	},
	{
		match:   "fl",
		replace: "ফ্ল",
	},
	{
		match:   "f",
		replace: "ফ",
	},
	{
		match:   "p",
		replace: "প",
	},
	{
		match:   "rri`",
		replace: "ৃ",
	},
	{
		match:   "rri",
		replace: "ৃ",
		exceptions: []exception{
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   notConsonant,
					},
				},
				thenReplace: "ঋ",
			},
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   punctuation,
					},
				},
				thenReplace: "ঋ",
			},
		},
	},
	{
		match:   "rrZ",
		replace: "রর‍্য",
	},
	{
		match:   "rry",
		replace: "রর‍্য",
	},
	{
		match:   "rZ",
		replace: "র‍্য",
		exceptions: []exception{
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   consonant,
					},
					{
						when:  prefix,
						is:    notExactly,
						value: "r",
					},
					{
						when:  prefix,
						is:    notExactly,
						value: "y",
					},
					{
						when:  prefix,
						is:    notExactly,
						value: "w",
					},
					{
						when:  prefix,
						is:    notExactly,
						value: "x",
					},
				},
				thenReplace: "্র্য",
			},
		},
	},
	{
		match:   "ry",
		replace: "র‍্য",
		exceptions: []exception{
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   consonant,
					},
					{
						when:  prefix,
						is:    notExactly,
						value: "r",
					},
					{
						when:  prefix,
						is:    notExactly,
						value: "y",
					},
					{
						when:  prefix,
						is:    notExactly,
						value: "w",
					},
					{
						when:  prefix,
						is:    notExactly,
						value: "x",
					},
				},
				thenReplace: "্র্য",
			},
		},
	},
	{
		match:   "rr",
		replace: "রর",
		exceptions: []exception{
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   notConsonant,
					},
					{
						when: suffix,
						is:   notVowel,
					},
					{
						when:  suffix,
						is:    notExactly,
						value: "r",
					},
					{
						when: suffix,
						is:   notPunctuation,
					},
				},
				thenReplace: "র্",
			},
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   consonant,
					},
					{
						when:  prefix,
						is:    notExactly,
						value: "r",
					},
				},
				thenReplace: "্রর",
			},
		},
	},
	{
		match:   "Rg",
		replace: bnRRA + "্গ",
	},
	{
		match:   "Rh",
		replace: bnRHA,
	},
	{
		match:   "R",
		replace: bnRRA,
	},
	{
		match:   "r",
		replace: "র",
		exceptions: []exception{
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   consonant,
					},
					{
						when:  prefix,
						is:    notExactly,
						value: "r",
					},
					{
						when:  prefix,
						is:    notExactly,
						value: "y",
					},
					{
						when:  prefix,
						is:    notExactly,
						value: "w",
					},
					{
						when:  prefix,
						is:    notExactly,
						value: "x",
					},
					{
						when:  prefix,
						is:    notExactly,
						value: "Z",
					},
				},
				thenReplace: "্র",
			},
		},
	},
	{
		match:   "shch",
		replace: "শ্ছ",
	},
	{
		match:   "ShTh",
		replace: "ষ্ঠ",
	},
	{
		match:   "Shph",
		replace: "ষ্ফ",
	},
	{
		match:   "Sch",
		replace: "শ্ছ",
	},
	{
		match:   "skl",
		replace: "স্ক্ল",
	},
	{
		match:   "skh",
		replace: "স্খ",
	},
	{
		match:   "sth",
		replace: "স্থ",
	},
	{
		match:   "sph",
		replace: "স্ফ",
	},
	{
		match:   "shc",
		replace: "শ্চ",
	},
	{
		match:   "sht",
		replace: "শ্ত",
	},
	{
		match:   "shn",
		replace: "শ্ন",
	},
	{
		match:   "shm",
		replace: "শ্ম",
	},
	{
		match:   "shl",
		replace: "শ্ল",
	},
	{
		match:   "Shk",
		replace: "ষ্ক",
	},
	{
		match:   "ShT",
		replace: "ষ্ট",
	},
	{
		match:   "ShN",
		replace: "ষ্ণ",
	},
	{
		match:   "Shp",
		replace: "ষ্প",
	},
	{
		match:   "Shf",
		replace: "ষ্ফ",
	},
	{
		match:   "Shm",
		replace: "ষ্ম",
	},
	{
		match:   "spl",
		replace: "স্প্ল",
	},
	{
		match:   "sk",
		replace: "স্ক",
	},
	{
		match:   "Sc",
		replace: "শ্চ",
	},
	{
		match:   "sT",
		replace: "স্ট",
	},
	{
		match:   "st",
		replace: "স্ত",
	},
	{
		match:   "sn",
		replace: "স্ন",
	},
	{
		match:   "sp",
		replace: "স্প",
	},
	{
		match:   "sf",
		replace: "স্ফ",
	},
	{
		match:   "sm",
		replace: "স্ম",
	},
	{
		match:   "sl",
		replace: "স্ল",
	},
	{
		match:   "sh",
		replace: "শ",
	},
	{
		match:   "Sc",
		replace: "শ্চ",
	},
	{
		match:   "St",
		replace: "শ্ত",
	},
	{
		match:   "Sn",
		replace: "শ্ন",
	},
	{
		match:   "Sm",
		replace: "শ্ম",
	},
	{
		match:   "Sl",
		replace: "শ্ল",
	},
	{
		match:   "Sh",
		replace: "ষ",
	},
	{
		match:   "s",
		replace: "স",
	},
	{
		match:   "S",
		replace: "শ",
	},
	{
		match:   "oo`",
		replace: "ু",
	},
	{
		match:   "oo",
		replace: "ু",
		exceptions: []exception{
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   notConsonant,
					},
					{
						when:  suffix,
						is:    notExactly,
						value: "`",
					},
				},
				thenReplace: "উ",
			},
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   punctuation,
					},
					{
						when:  suffix,
						is:    notExactly,
						value: "`",
					},
				},
				thenReplace: "উ",
			},
		},
	},
	{
		match:   "o`",
		replace: "",
	},
	{
		match:   "oZ",
		replace: "অ্য",
	},
	{
		match:   "o",
		replace: "",
		exceptions: []exception{
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   vowel,
					},
					{
						when:  prefix,
						is:    notExactly,
						value: "o",
					},
				},
				thenReplace: "ও",
			},
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   vowel,
					},
					{
						when:  prefix,
						is:    exactly,
						value: "o",
					},
				},
				thenReplace: "অ",
			},
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   punctuation,
					},
				},
				thenReplace: "অ",
			},
		},
	},
	{
		match:   "tth",
		replace: "ত্থ",
	},
	{
		match:   "t``",
		replace: "ৎ",
	},
	{
		match:   "TT",
		replace: "ট্ট",
	},
	{
		match:   "Tm",
		replace: "ট্ম",
	},
	{
		match:   "Th",
		replace: "ঠ",
	},
	{
		match:   "tn",
		replace: "ত্ন",
	},
	{
		match:   "tm",
		replace: "ত্ম",
	},
	{
		match:   "th",
		replace: "থ",
	},
	{
		match:   "tt",
		replace: "ত্ত",
	},
	{
		match:   "T",
		replace: "ট",
	},
	{
		match:   "t",
		replace: "ত",
	},
	{
		match:   "aZ",
		replace: "অ্যা",
	},
	{
		match:   "AZ",
		replace: "অ্যা",
	},
	{
		match:   "a`",
		replace: "া",
	},
	{
		match:   "A`",
		replace: "া",
	},
	{
		match:   "i`",
		replace: "ি",
	},
	{
		match:   "i",
		replace: "ি",
		exceptions: []exception{
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   notConsonant,
					},
					{
						when:  suffix,
						is:    notExactly,
						value: "`",
					},
				},
				thenReplace: "ই",
			},
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   punctuation,
					},
					{
						when:  suffix,
						is:    notExactly,
						value: "`",
					},
				},
				thenReplace: "ই",
			},
		},
	},
	{
		match:   "I`",
		replace: "ী",
	},
	{
		match:   "I",
		replace: "ী",
		exceptions: []exception{
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   notConsonant,
					},
					{
						when:  suffix,
						is:    notExactly,
						value: "`",
					},
				},
				thenReplace: "ঈ",
			},
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   punctuation,
					},
					{
						when:  suffix,
						is:    notExactly,
						value: "`",
					},
				},
				thenReplace: "ঈ",
			},
		},
	},
	{
		match:   "u`",
		replace: "ু",
	},
	{
		match:   "u",
		replace: "ু",
		exceptions: []exception{
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   notConsonant,
					},
					{
						when:  suffix,
						is:    notExactly,
						value: "`",
					},
				},
				thenReplace: "উ",
			},
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   punctuation,
					},
					{
						when:  suffix,
						is:    notExactly,
						value: "`",
					},
				},
				thenReplace: "উ",
			},
		},
	},
	{
		match:   "U`",
		replace: "ূ",
	},
	{
		match:   "U",
		replace: "ূ",
		exceptions: []exception{
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   notConsonant,
					},
					{
						when:  suffix,
						is:    notExactly,
						value: "`",
					},
				},
				thenReplace: "ঊ",
			},
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   punctuation,
					},
					{
						when:  suffix,
						is:    notExactly,
						value: "`",
					},
				},
				thenReplace: "ঊ",
			},
		},
	},
	{
		match:   "ee`",
		replace: "ী",
	},
	{
		match:   "ee",
		replace: "ী",
		exceptions: []exception{
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   notConsonant,
					},
					{
						when:  suffix,
						is:    notExactly,
						value: "`",
					},
				},
				thenReplace: "ঈ",
			},
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   punctuation,
					},
					{
						when:  suffix,
						is:    notExactly,
						value: "`",
					},
				},
				thenReplace: "ঈ",
			},
		},
	},
	{
		match:   "e`",
		replace: "ে",
	},
	{
		match:   "e",
		replace: "ে",
		exceptions: []exception{
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   notConsonant,
					},
					{
						when:  suffix,
						is:    notExactly,
						value: "`",
					},
				},
				thenReplace: "এ",
			},
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   punctuation,
					},
					{
						when:  suffix,
						is:    notExactly,
						value: "`",
					},
				},
				thenReplace: "এ",
			},
		},
	},
	{
		match:   "z",
		replace: "য",
	},
	{
		match:   "Z",
		replace: "্য",
	},
	{
		match:   "y",
		replace: "্য",
		exceptions: []exception{
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   notConsonant,
					},
					{
						when: prefix,
						is:   notPunctuation,
					},
				},
				thenReplace: bnYYA,
			},
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   punctuation,
					},
				},
				thenReplace: "ই" + bnYYA,
			},
		},
	},
	{
		match:   "Y",
		replace: bnYYA,
	},
	{
		match:   "q",
		replace: "ক",
	},
	{
		match:   "w",
		replace: "ও",
		exceptions: []exception{
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   punctuation,
					},
					{
						when: suffix,
						is:   vowel,
					},
				},
				thenReplace: "ও" + bnYYA,
			},
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   consonant,
					},
				},
				thenReplace: "্ব",
			},
		},
	},
	{
		match:   "x",
		replace: "ক্স",
		exceptions: []exception{
			{
				ifAllMatch: []matchCondition{
					{
						when: prefix,
						is:   punctuation,
					},
				},
				thenReplace: "এক্স",
			},
		},
	},
	{
		match:   ":`",
		replace: ":",
	},
	{
		match:   ":",
		replace: "ঃ",
	},
	{
		match:   "^`",
		replace: "^",
	},
	{
		match:   "^",
		replace: "ঁ",
	},
	{
		match:   ",,",
		replace: "্‌",
	},
	{
		match:   ",",
		replace: ",",
	},
	{
		match:   "$",
		replace: "৳",
	},
	{
		match:   "`",
		replace: "",
	},
}
