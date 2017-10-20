package pkg

/**
 * @author Phil Teare
 * using wikipedia data
 */

type IsoLang struct {
	EnglishName string
	NativeName  string
}

var IsoLangs = map[string]IsoLang{
	"ab": {
		EnglishName: "Abkhaz",
		NativeName:  "аҧсуа",
	},
	"aa": {
		EnglishName: "Afar",
		NativeName:  "Afaraf",
	},
	"af": {
		EnglishName: "Afrikaans",
		NativeName:  "Afrikaans",
	},
	"ak": {
		EnglishName: "Akan",
		NativeName:  "Akan",
	},
	"sq": {
		EnglishName: "Albanian",
		NativeName:  "Shqip",
	},
	"am": {
		EnglishName: "Amharic",
		NativeName:  "አማርኛ",
	},
	"ar": {
		EnglishName: "Arabic",
		NativeName:  "العربية",
	},
	"an": {
		EnglishName: "Aragonese",
		NativeName:  "Aragonés",
	},
	"hy": {
		EnglishName: "Armenian",
		NativeName:  "Հայերեն",
	},
	"as": {
		EnglishName: "Assamese",
		NativeName:  "অসমীয়া",
	},
	"av": {
		EnglishName: "Avaric",
		NativeName:  "авар мацӀ, магӀарул мацӀ",
	},
	"ae": {
		EnglishName: "Avestan",
		NativeName:  "avesta",
	},
	"ay": {
		EnglishName: "Aymara",
		NativeName:  "aymar aru",
	},
	"az": {
		EnglishName: "Azerbaijani",
		NativeName:  "azərbaycan dili",
	},
	"bm": {
		EnglishName: "Bambara",
		NativeName:  "bamanankan",
	},
	"ba": {
		EnglishName: "Bashkir",
		NativeName:  "башҡорт теле",
	},
	"eu": {
		EnglishName: "Basque",
		NativeName:  "euskara, euskera",
	},
	"be": {
		EnglishName: "Belarusian",
		NativeName:  "Беларуская",
	},
	"bn": {
		EnglishName: "Bengali",
		NativeName:  "বাংলা",
	},
	"bh": {
		EnglishName: "Bihari",
		NativeName:  "भोजपुरी",
	},
	"bi": {
		EnglishName: "Bislama",
		NativeName:  "Bislama",
	},
	"bs": {
		EnglishName: "Bosnian",
		NativeName:  "bosanski jezik",
	},
	"br": {
		EnglishName: "Breton",
		NativeName:  "brezhoneg",
	},
	"bg": {
		EnglishName: "Bulgarian",
		NativeName:  "български език",
	},
	"my": {
		EnglishName: "Burmese",
		NativeName:  "ဗမာစာ",
	},
	"ca": {
		EnglishName: "Catalan; Valencian",
		NativeName:  "Català",
	},
	"ch": {
		EnglishName: "Chamorro",
		NativeName:  "Chamoru",
	},
	"ce": {
		EnglishName: "Chechen",
		NativeName:  "нохчийн мотт",
	},
	"ny": {
		EnglishName: "Chichewa; Chewa; Nyanja",
		NativeName:  "chiCheŵa, chinyanja",
	},
	"zh": {
		EnglishName: "Chinese",
		NativeName:  "中文 (Zhōngwén), 汉语, 漢語",
	},
	"cv": {
		EnglishName: "Chuvash",
		NativeName:  "чӑваш чӗлхи",
	},
	"kw": {
		EnglishName: "Cornish",
		NativeName:  "Kernewek",
	},
	"co": {
		EnglishName: "Corsican",
		NativeName:  "corsu, lingua corsa",
	},
	"cr": {
		EnglishName: "Cree",
		NativeName:  "ᓀᐦᐃᔭᐍᐏᐣ",
	},
	"hr": {
		EnglishName: "Croatian",
		NativeName:  "hrvatski",
	},
	"cs": {
		EnglishName: "Czech",
		NativeName:  "česky, čeština",
	},
	"da": {
		EnglishName: "Danish",
		NativeName:  "dansk",
	},
	"dv": {
		EnglishName: "Divehi; Dhivehi; Maldivian;",
		NativeName:  "ދިވެހި",
	},
	"nl": {
		EnglishName: "Dutch",
		NativeName:  "Nederlands, Vlaams",
	},
	"en": {
		EnglishName: "English",
		NativeName:  "English",
	},
	"eo": {
		EnglishName: "Esperanto",
		NativeName:  "Esperanto",
	},
	"et": {
		EnglishName: "Estonian",
		NativeName:  "eesti, eesti keel",
	},
	"ee": {
		EnglishName: "Ewe",
		NativeName:  "Eʋegbe",
	},
	"fo": {
		EnglishName: "Faroese",
		NativeName:  "føroyskt",
	},
	"fj": {
		EnglishName: "Fijian",
		NativeName:  "vosa Vakaviti",
	},
	"fi": {
		EnglishName: "Finnish",
		NativeName:  "suomi, suomen kieli",
	},
	"fr": {
		EnglishName: "French",
		NativeName:  "français, langue française",
	},
	"ff": {
		EnglishName: "Fula; Fulah; Pulaar; Pular",
		NativeName:  "Fulfulde, Pulaar, Pular",
	},
	"gl": {
		EnglishName: "Galician",
		NativeName:  "Galego",
	},
	"ka": {
		EnglishName: "Georgian",
		NativeName:  "ქართული",
	},
	"de": {
		EnglishName: "German",
		NativeName:  "Deutsch",
	},
	"el": {
		EnglishName: "Greek, Modern",
		NativeName:  "Ελληνικά",
	},
	"gn": {
		EnglishName: "Guaraní",
		NativeName:  "Avañeẽ",
	},
	"gu": {
		EnglishName: "Gujarati",
		NativeName:  "ગુજરાતી",
	},
	"ht": {
		EnglishName: "Haitian; Haitian Creole",
		NativeName:  "Kreyòl ayisyen",
	},
	"ha": {
		EnglishName: "Hausa",
		NativeName:  "Hausa, هَوُسَ",
	},
	"he": {
		EnglishName: "Hebrew (modern)",
		NativeName:  "עברית",
	},
	"hz": {
		EnglishName: "Herero",
		NativeName:  "Otjiherero",
	},
	"hi": {
		EnglishName: "Hindi",
		NativeName:  "हिन्दी, हिंदी",
	},
	"ho": {
		EnglishName: "Hiri Motu",
		NativeName:  "Hiri Motu",
	},
	"hu": {
		EnglishName: "Hungarian",
		NativeName:  "Magyar",
	},
	"ia": {
		EnglishName: "Interlingua",
		NativeName:  "Interlingua",
	},
	"id": {
		EnglishName: "Indonesian",
		NativeName:  "Bahasa Indonesia",
	},
	"ie": {
		EnglishName: "Interlingue",
		NativeName:  "Originally called Occidental; then Interlingue after WWII",
	},
	"ga": {
		EnglishName: "Irish",
		NativeName:  "Gaeilge",
	},
	"ig": {
		EnglishName: "Igbo",
		NativeName:  "Asụsụ Igbo",
	},
	"ik": {
		EnglishName: "Inupiaq",
		NativeName:  "Iñupiaq, Iñupiatun",
	},
	"io": {
		EnglishName: "Ido",
		NativeName:  "Ido",
	},
	"is": {
		EnglishName: "Icelandic",
		NativeName:  "Íslenska",
	},
	"it": {
		EnglishName: "Italian",
		NativeName:  "Italiano",
	},
	"iu": {
		EnglishName: "Inuktitut",
		NativeName:  "ᐃᓄᒃᑎᑐᑦ",
	},
	"ja": {
		EnglishName: "Japanese",
		NativeName:  "日本語 (にほんご／にっぽんご)",
	},
	"jv": {
		EnglishName: "Javanese",
		NativeName:  "basa Jawa",
	},
	"kl": {
		EnglishName: "Kalaallisut, Greenlandic",
		NativeName:  "kalaallisut, kalaallit oqaasii",
	},
	"kn": {
		EnglishName: "Kannada",
		NativeName:  "ಕನ್ನಡ",
	},
	"kr": {
		EnglishName: "Kanuri",
		NativeName:  "Kanuri",
	},
	"ks": {
		EnglishName: "Kashmiri",
		NativeName:  "कश्मीरी, كشميري‎",
	},
	"kk": {
		EnglishName: "Kazakh",
		NativeName:  "Қазақ тілі",
	},
	"km": {
		EnglishName: "Khmer",
		NativeName:  "ភាសាខ្មែរ",
	},
	"ki": {
		EnglishName: "Kikuyu, Gikuyu",
		NativeName:  "Gĩkũyũ",
	},
	"rw": {
		EnglishName: "Kinyarwanda",
		NativeName:  "Ikinyarwanda",
	},
	"ky": {
		EnglishName: "Kirghiz, Kyrgyz",
		NativeName:  "кыргыз тили",
	},
	"kv": {
		EnglishName: "Komi",
		NativeName:  "коми кыв",
	},
	"kg": {
		EnglishName: "Kongo",
		NativeName:  "KiKongo",
	},
	"ko": {
		EnglishName: "Korean",
		NativeName:  "한국어 (韓國語), 조선말 (朝鮮語)",
	},
	"ku": {
		EnglishName: "Kurdish",
		NativeName:  "Kurdî, كوردی‎",
	},
	"kj": {
		EnglishName: "Kwanyama, Kuanyama",
		NativeName:  "Kuanyama",
	},
	"la": {
		EnglishName: "Latin",
		NativeName:  "latine, lingua latina",
	},
	"lb": {
		EnglishName: "Luxembourgish, Letzeburgesch",
		NativeName:  "Lëtzebuergesch",
	},
	"lg": {
		EnglishName: "Luganda",
		NativeName:  "Luganda",
	},
	"li": {
		EnglishName: "Limburgish, Limburgan, Limburger",
		NativeName:  "Limburgs",
	},
	"ln": {
		EnglishName: "Lingala",
		NativeName:  "Lingála",
	},
	"lo": {
		EnglishName: "Lao",
		NativeName:  "ພາສາລາວ",
	},
	"lt": {
		EnglishName: "Lithuanian",
		NativeName:  "lietuvių kalba",
	},
	"lu": {
		EnglishName: "Luba-Katanga",
		NativeName:  "",
	},
	"lv": {
		EnglishName: "Latvian",
		NativeName:  "latviešu valoda",
	},
	"gv": {
		EnglishName: "Manx",
		NativeName:  "Gaelg, Gailck",
	},
	"mk": {
		EnglishName: "Macedonian",
		NativeName:  "македонски јазик",
	},
	"mg": {
		EnglishName: "Malagasy",
		NativeName:  "Malagasy fiteny",
	},
	"ms": {
		EnglishName: "Malay",
		NativeName:  "bahasa Melayu, بهاس ملايو‎",
	},
	"ml": {
		EnglishName: "Malayalam",
		NativeName:  "മലയാളം",
	},
	"mt": {
		EnglishName: "Maltese",
		NativeName:  "Malti",
	},
	"mi": {
		EnglishName: "Māori",
		NativeName:  "te reo Māori",
	},
	"mr": {
		EnglishName: "Marathi (Marāṭhī)",
		NativeName:  "मराठी",
	},
	"mh": {
		EnglishName: "Marshallese",
		NativeName:  "Kajin M̧ajeļ",
	},
	"mn": {
		EnglishName: "Mongolian",
		NativeName:  "монгол",
	},
	"na": {
		EnglishName: "Nauru",
		NativeName:  "Ekakairũ Naoero",
	},
	"nv": {
		EnglishName: "Navajo, Navaho",
		NativeName:  "Diné bizaad, Dinékʼehǰí",
	},
	"nb": {
		EnglishName: "Norwegian Bokmål",
		NativeName:  "Norsk bokmål",
	},
	"nd": {
		EnglishName: "North Ndebele",
		NativeName:  "isiNdebele",
	},
	"ne": {
		EnglishName: "Nepali",
		NativeName:  "नेपाली",
	},
	"ng": {
		EnglishName: "Ndonga",
		NativeName:  "Owambo",
	},
	"nn": {
		EnglishName: "Norwegian Nynorsk",
		NativeName:  "Norsk nynorsk",
	},
	"no": {
		EnglishName: "Norwegian",
		NativeName:  "Norsk",
	},
	"ii": {
		EnglishName: "Nuosu",
		NativeName:  "ꆈꌠ꒿ Nuosuhxop",
	},
	"nr": {
		EnglishName: "South Ndebele",
		NativeName:  "isiNdebele",
	},
	"oc": {
		EnglishName: "Occitan",
		NativeName:  "Occitan",
	},
	"oj": {
		EnglishName: "Ojibwe, Ojibwa",
		NativeName:  "ᐊᓂᔑᓈᐯᒧᐎᓐ",
	},
	"cu": {
		EnglishName: "Old Church Slavonic, Church Slavic, Church Slavonic, Old Bulgarian, Old Slavonic",
		NativeName:  "ѩзыкъ словѣньскъ",
	},
	"om": {
		EnglishName: "Oromo",
		NativeName:  "Afaan Oromoo",
	},
	"or": {
		EnglishName: "Oriya",
		NativeName:  "ଓଡ଼ିଆ",
	},
	"os": {
		EnglishName: "Ossetian, Ossetic",
		NativeName:  "ирон æвзаг",
	},
	"pa": {
		EnglishName: "Panjabi, Punjabi",
		NativeName:  "ਪੰਜਾਬੀ, پنجابی‎",
	},
	"pi": {
		EnglishName: "Pāli",
		NativeName:  "पाऴि",
	},
	"fa": {
		EnglishName: "Persian",
		NativeName:  "فارسی",
	},
	"pl": {
		EnglishName: "Polish",
		NativeName:  "polski",
	},
	"ps": {
		EnglishName: "Pashto, Pushto",
		NativeName:  "پښتو",
	},
	"pt": {
		EnglishName: "Portuguese",
		NativeName:  "Português",
	},
	"qu": {
		EnglishName: "Quechua",
		NativeName:  "Runa Simi, Kichwa",
	},
	"rm": {
		EnglishName: "Romansh",
		NativeName:  "rumantsch grischun",
	},
	"rn": {
		EnglishName: "Kirundi",
		NativeName:  "kiRundi",
	},
	"ro": {
		EnglishName: "Romanian, Moldavian, Moldovan",
		NativeName:  "română",
	},
	"ru": {
		EnglishName: "Russian",
		NativeName:  "русский язык",
	},
	"sa": {
		EnglishName: "Sanskrit (Saṁskṛta)",
		NativeName:  "संस्कृतम्",
	},
	"sc": {
		EnglishName: "Sardinian",
		NativeName:  "sardu",
	},
	"sd": {
		EnglishName: "Sindhi",
		NativeName:  "सिन्धी, سنڌي، سندھی‎",
	},
	"se": {
		EnglishName: "Northern Sami",
		NativeName:  "Davvisámegiella",
	},
	"sm": {
		EnglishName: "Samoan",
		NativeName:  "gagana faa Samoa",
	},
	"sg": {
		EnglishName: "Sango",
		NativeName:  "yângâ tî sängö",
	},
	"sr": {
		EnglishName: "Serbian",
		NativeName:  "српски језик",
	},
	"gd": {
		EnglishName: "Scottish Gaelic; Gaelic",
		NativeName:  "Gàidhlig",
	},
	"sn": {
		EnglishName: "Shona",
		NativeName:  "chiShona",
	},
	"si": {
		EnglishName: "Sinhala, Sinhalese",
		NativeName:  "සිංහල",
	},
	"sk": {
		EnglishName: "Slovak",
		NativeName:  "slovenčina",
	},
	"sl": {
		EnglishName: "Slovene",
		NativeName:  "slovenščina",
	},
	"so": {
		EnglishName: "Somali",
		NativeName:  "Soomaaliga, af Soomaali",
	},
	"st": {
		EnglishName: "Southern Sotho",
		NativeName:  "Sesotho",
	},
	"es": {
		EnglishName: "Spanish; Castilian",
		NativeName:  "español, castellano",
	},
	"su": {
		EnglishName: "Sundanese",
		NativeName:  "Basa Sunda",
	},
	"sw": {
		EnglishName: "Swahili",
		NativeName:  "Kiswahili",
	},
	"ss": {
		EnglishName: "Swati",
		NativeName:  "SiSwati",
	},
	"sv": {
		EnglishName: "Swedish",
		NativeName:  "svenska",
	},
	"ta": {
		EnglishName: "Tamil",
		NativeName:  "தமிழ்",
	},
	"te": {
		EnglishName: "Telugu",
		NativeName:  "తెలుగు",
	},
	"tg": {
		EnglishName: "Tajik",
		NativeName:  "тоҷикӣ, toğikī, تاجیکی‎",
	},
	"th": {
		EnglishName: "Thai",
		NativeName:  "ไทย",
	},
	"ti": {
		EnglishName: "Tigrinya",
		NativeName:  "ትግርኛ",
	},
	"bo": {
		EnglishName: "Tibetan Standard, Tibetan, Central",
		NativeName:  "བོད་ཡིག",
	},
	"tk": {
		EnglishName: "Turkmen",
		NativeName:  "Türkmen, Түркмен",
	},
	"tl": {
		EnglishName: "Tagalog",
		NativeName:  "Wikang Tagalog, ᜏᜒᜃᜅ᜔ ᜆᜄᜎᜓᜄ᜔",
	},
	"tn": {
		EnglishName: "Tswana",
		NativeName:  "Setswana",
	},
	"to": {
		EnglishName: "Tonga (Tonga Islands)",
		NativeName:  "faka Tonga",
	},
	"tr": {
		EnglishName: "Turkish",
		NativeName:  "Türkçe",
	},
	"ts": {
		EnglishName: "Tsonga",
		NativeName:  "Xitsonga",
	},
	"tt": {
		EnglishName: "Tatar",
		NativeName:  "татарча, tatarça, تاتارچا‎",
	},
	"tw": {
		EnglishName: "Twi",
		NativeName:  "Twi",
	},
	"ty": {
		EnglishName: "Tahitian",
		NativeName:  "Reo Tahiti",
	},
	"ug": {
		EnglishName: "Uighur, Uyghur",
		NativeName:  "Uyƣurqə, ئۇيغۇرچە‎",
	},
	"uk": {
		EnglishName: "Ukrainian",
		NativeName:  "українська",
	},
	"ur": {
		EnglishName: "Urdu",
		NativeName:  "اردو",
	},
	"uz": {
		EnglishName: "Uzbek",
		NativeName:  "zbek, Ўзбек, أۇزبېك‎",
	},
	"ve": {
		EnglishName: "Venda",
		NativeName:  "Tshivenḓa",
	},
	"vi": {
		EnglishName: "Vietnamese",
		NativeName:  "Tiếng Việt",
	},
	"vo": {
		EnglishName: "Volapük",
		NativeName:  "Volapük",
	},
	"wa": {
		EnglishName: "Walloon",
		NativeName:  "Walon",
	},
	"cy": {
		EnglishName: "Welsh",
		NativeName:  "Cymraeg",
	},
	"wo": {
		EnglishName: "Wolof",
		NativeName:  "Wollof",
	},
	"fy": {
		EnglishName: "Western Frisian",
		NativeName:  "Frysk",
	},
	"xh": {
		EnglishName: "Xhosa",
		NativeName:  "isiXhosa",
	},
	"yi": {
		EnglishName: "Yiddish",
		NativeName:  "ייִדיש",
	},
	"yo": {
		EnglishName: "Yoruba",
		NativeName:  "Yorùbá",
	},
	"za": {
		EnglishName: "Zhuang, Chuang",
		NativeName:  "Saɯ cueŋƅ, Saw cuengh",
	},
}
