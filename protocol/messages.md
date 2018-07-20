# Messages

## Status [STA]

+ Positional Parameters
	+ Code (StatusCode)
	+ Description (string)

+ Flags
	+ FC, TL, TO, PR, FM, FB, I4, I6; BASE $ 5.3.1. STA (BASE v1.0.3)
	+ RF; EXT § 3.10 RF - Referrer notification (EXT v1.0.8)
	+ QP; EXT § 3.11 QP - Upload queue notification (EXT v1.0.8)
	+ FC, TO, RC; EXT § 3.27 ASCH - Extended searching capability (EXT v1.0.8)

## Supported Features [SUP]

+ Positional Parameters
	+ FeatureOps ([]FeatureOp)

## Session ID [SID]

+ Positional Parameters
	+ SID (base32)

## Info [INF]

+ Named Parameters
	+ ID (base32) - The CID of the client. Mandatory for C-C connections. Specified in BASE.
	+ PD (base32) - The PID of the client. Hubs must check that the hash(PID) == CID and then discard the field before broadcasting it to other clients. Must not be sent in C-C connections. Specified in BASE.

	+ I4 (ip) - IPv4 address without port. A zero address (0.0.0.0) means that the server should replace it with the real IP of the client. Hubs must check that a specified address corresponds to what the client is connecting from to avoid DoS attacks and only allow trusted clients to specify a different address. Clients should use the zero address when connecting, but may opt not to do so at the user’s discretion. Specified in BASE.
	+ I6 (ip) - IPv6 address without port. A zero address (::) means that the server should replace it with the IP of the client. Specified in BASE.

	+ U4 (int) - The Client UDP port. Specified in BASE.
	+ U6 (int) - Same as U4, but for IPv6. Specified in BASE.

	+ SS (int) - Share size in bytes. Specified in BASE.
	+ SF (int) - Number of shared files. Specified in BASE.

	+ VE (string) - Client identification, version (client-specific, a short identifier then a dotted version number is recommended). Specified in BASE.

	+ US (int) - Maximum upload speed, bytes/second. Specified in BASE.
	+ DS (int) - Maximum downloads speed, bytes/second. Specified in BASE.

	+ SL (int) - Maximum simultaneous upload connections (slots). Specified in BASE.

	+ AS (int) - Automatic slot allocator speed limit, bytes/sec. The client keeps opening slots as long as its total upload speed doesn’t exceed this value. Specified in BASE.
	+ AM (int) - Minimum simultaneous upload connections in automatic slot manager mode. Specified in BASE.

	+ EM (string) - E-mail address. Specified in BASE.
	+ NI (string) - Nickname (or hub name). The hub must ensure that this is unique in the hub up to case-sensitivity. Valid are all characters in the Unicode character set with code point above 32, although hubs may limit this further as they like with an appropriate error message. Specified in BASE.
	+ DE (string) - Description. Valid are all characters in the Unicode character set with code point equal to or greater than 32. Specified in BASE.

	+ HN (int) - Hubs where user is a normal user and in NORMAL state. Specified in BASE.
	+ HR (int) - Hubs where user is registered (had to supply password) and in NORMAL state. Specified in BASE.
	+ HO (int) - Hubs where user is op and in NORMAL state. Specified in BASE.

	+ TO (string) - Token, as received in RCM/CTM, when establishing a C-C connection. Specified in BASE.

	+ CT (int) - Client (user) type, 1=bot, 2=registered user, 4=operator, 8=super user, 16=hub owner, 32=hub (used when the hub sends an INF about itself). Multiple types are specified by adding the numbers together. Specified in BASE.
	+ AW (int) - 1=Away, 2=Extended away, not interested in hub chat (hubs may skip sending broadcast type MSG commands to clients with this flag). Specified in BASE.

	+ SU ([]string) - Comma-separated list of feature FOURCC’s. This notifies other clients of extended capabilities of the connecting client. Specified in BASE.
	+ RF (string) - URL of referrer (hub in case of redirect, web page). Specified in BASE.

+ Flags
	+ HH, WS, NE, OW, UC, SS, SF, MS, XS, ML, XL, MU, MR, MO, XU, XR, XO, MC, UP; EXT $ 3.4 PING - Pinger extension (EXT v1.0.8)
	+ LC; EXT § 3.13 LC - Locale specification (EXT v1.0.8)
	+ KP; EXT § 3.16 KEYP - Certificate substitution protection in conjunction with ADCS (EXT v1.0.8)
	+ FO; EXT § 3.21 FO - Failover hub addresses (EXT v1.0.8)
	+ FS; EXT § 3.22 FS - Free slots in client (EXT v1.0.8)
	+ AP; EXT § 3.24 Application and version separation in INF (EXT v1.0.8)
	+ RP; EXT § 3.32 RDEX - Redirects Extended (EXT v1.0.8)

## Message [MSG]

+ Positional Parameters
	+ Text (string)

+ Named Parameters
	+ PM (base32)
	+ ME (int)

+ Flags
	+ TS, EXT § 3.5 TS - Timestamp in MSG (EXT v1.0.8)

## Search [SCH]

+ Positional Parameters
	+ SearchTerms ([]SearchTerm)

+ Named Parameters
	+ LE (int) - Smaller (less) than or equal size in bytes. Specified in BASE.
	+ GE (int) - Larger (greater) than or equal size in bytes. Specified in BASE.
	+ EQ (int) - Exact size in bytes. Specified in BASE.
	+ TO (string) - Token, string. Used by the client to tell one search from the other. If present, the responding client must copy this field to each search result. Specified in BASE.
	+ TY (int) - File type, to be chosen from the following (none specified = any type): 1 = File, 2 = Directory. Specified in BASE.
	+ TR (base32) - Tiger tree Hash root, encoded with base32. Specified in EXT § 3.1 TIGR - Tiger tree hash support (EXT v1.0.8).
	+ TD (int) - Tree depth, index of the highest level of tree data available, root-only = 0, first level (2 leaves) = 1, second level = 2, etc… Specified in EXT § 3.1 TIGR - Tiger tree hash support (EXT v1.0.8).

+ Flags
	+ KY; EXT § 3.17. SUDP - Encrypting UDP traffic (EXT v1.0.8)
	+ GR, RX; EXT § 3.20 SEGA - Grouping of file extensions in SCH (EXT v1.0.8)
	+ MT, PP, OT, NT, MR, PA, RE; EXT § 3.27 ASCH - Extended searching capability (EXT v1.0.8)

## Result [RES]

+ Named Parameters
	+ FN (string, required)
	+ SI (int, required)
	+ SL (int)
	+ TO (string, required)

	+ TR (base32) - Tiger tree Hash root, encoded with base32. Specified in EXT § 3.1 TIGR - Tiger tree hash support (EXT v1.0.8).
	+ TD (int) - Tree depth, index of the highest level of tree data available, root-only = 0, first level (2 leaves) = 1, second level = 2, etc… Specified in EXT § 3.1 TIGR - Tiger tree hash support (EXT v1.0.8).

+ Flags
	+ FI, FO, DA; EXT § 3.27 ASCH - Extended searching capability (EXT v1.0.8)

## Connect To Me [CTM]

+ Positional Parameters
	+ Protocol (string)
	+ Port (string)
	+ Token (string)

## Reverse Connect To Me [RCM]

+ Positional Parameters
	+ Protocol (string)
	+ Token (string)

+ Flags
	+ KY?; EXT § 3.17. SUDP - Encrypting UDP traffic (EXT v1.0.8)

## Get Password [GPA]

+ Positional Parameters
	+ Data (base32)

## Password [PAS]

+ Positional Parameters
	+ Password (base32)

## Quit [QUI]

+ Positional Parameters
	+ SID (base32)

+ Named Parameters
	+ ID (base32)
	+ TL (int)
	+ MS (string)
	+ RD (string)
	+ DI (string)

+ Flags
	+ RX, PT; EXT § 3.32 RDEX - Redirects Extended (EXT v1.0.8)

## Get File [GET]

+ Positional Parameters
	+ Namespace (string) - Specified in BASE. Known values. file, list; BASE $ 5.3.13. GET (BASE v1.0.3). tthl; EXT $ 3.1 TIGR - Tiger tree hash support (EXT v1.0.8). blom; EXT $ 3.8 BLOM - Bloom filter (EXT v1.0.8).
	+ Identifer (string)
	+ StartPos (int)
	+ Bytes (int)

+ Named Parameters
	+ RE (int)

+ Flags
	+ ZL; EXT § 3.3. ZLIB - Compressed communication (EXT v1.0.8)
	+ BK, BH; EXT $ 3.8 BLOM - Bloom filter (EXT v1.0.8)
	+ DB; EXT § 3.31 Downloaded progress report for uploaders in GET (EXT v1.0.8)

## Get File Info [GFI]

+ Positional Parameters
	+ Namespace (string) - Specified in BASE. Known values. file, list; BASE $ 5.3.13. GET (BASE v1.0.3). tthl; EXT $ 3.1 TIGR - Tiger tree hash support (EXT v1.0.8). blom; EXT $ 3.8 BLOM - Bloom filter (EXT v1.0.8).
	+ Identifer (string)

## Send File [SND]

+ Positional Parameters
	+ Namespace (string) - Specified in BASE. Known values. file, list; BASE $ 5.3.13. GET (BASE v1.0.3). tthl; EXT $ 3.1 TIGR - Tiger tree hash support (EXT v1.0.8). blom; EXT $ 3.8 BLOM - Bloom filter (EXT v1.0.8).
	+ Identifer (string)
	+ StartPos (int)
	+ Bytes (int)

+ Flags
	+ ZL; EXT § 3.3. ZLIB - Compressed communication (EXT v1.0.8)
