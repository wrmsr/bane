package irc

type Message struct {
	Name       string
	Parameters string
	Replies    []string
}

var Messages = []Message{
	{
		"PASS",
		"<password>",
		[]string{
			"ERR_NEEDMOREPARAMS",
			"ERR_ALREADYREGISTRED",
		},
	},
	{
		"NICK",
		"<nickname>",
		[]string{
			"ERR_NONICKNAMEGIVEN",
			"ERR_ERRONEUSNICKNAME",
			"ERR_NICKNAMEINUSE",
			"ERR_NICKCOLLISION",
			"ERR_UNAVAILRESOURCE",
			"ERR_RESTRICTED",
		},
	},

	{
		"USER",
		"<user> <mode> <unused> <realname>",
		[]string{
			"ERR_NEEDMOREPARAMS",
			"ERR_ALREADYREGISTRED",
		},
	},

	{
		"OPER",
		"<name> <password>",
		[]string{
			"ERR_NEEDMOREPARAMS",
			"RPL_YOUREOPER",
			"ERR_NOOPERHOST",
			"ERR_PASSWDMISMATCH",
		},
	},

	{
		"MODE",
		"(<nickname> *( ( \" + \" / \" - \" ) *( \"i\" / \"w\" / \"o\" / \"O\" / \"r\" ) )) | " + // nick
			"(<channel> *( ( \"-\" / \"+\" ) *<modes> *<modeparams> ))", // chan
		[]string{
			// nick
			"ERR_NEEDMOREPARAMS",
			"ERR_USERSDONTMATCH",
			"ERR_UMODEUNKNOWNFLAG",
			"RPL_UMODEIS",
			// chan
			"ERR_NEEDMOREPARAMS",
			"ERR_KEYSET",
			"ERR_NOCHANMODES",
			"ERR_CHANOPRIVSNEEDED",
			"ERR_USERNOTINCHANNEL",
			"ERR_UNKNOWNMODE",
			"RPL_CHANNELMODEIS",
			"RPL_BANLIST",
			"RPL_ENDOFBANLIST",
			"RPL_EXCEPTLIST",
			"RPL_ENDOFEXCEPTLIST",
			"RPL_INVITELIST",
			"RPL_ENDOFINVITELIST",
			"RPL_UNIQOPIS",
		},
	},

	{
		"SERVICE",
		"<nickname> <reserved> <distribution> <type> <reserved> <info>",
		[]string{
			"ERR_ALREADYREGISTRED",
			"ERR_NEEDMOREPARAMS",
			"ERR_ERRONEUSNICKNAME",
			"RPL_YOURESERVICE",
			"RPL_YOURHOST",
			"RPL_MYINFO",
		},
	},

	{
		"QUIT",
		"[ <Quit Message> ]",

		nil,
	},

	{
		"SQUIT",
		"<server> <comment>",
		[]string{
			"ERR_NOPRIVILEGES",
			"ERR_NOSUCHSERVER",
			"ERR_NEEDMOREPARAMS",
		},
	},

	{
		"JOIN",
		"( <channel> *( \",\" <channel> ) [ <key> *( \",\" <key> ) ] ) / \"0\"",
		[]string{
			"ERR_NEEDMOREPARAMS",
			"ERR_BANNEDFROMCHAN",
			"ERR_INVITEONLYCHAN",
			"ERR_BADCHANNELKEY",
			"ERR_CHANNELISFULL",
			"ERR_BADCHANMASK",
			"ERR_NOSUCHCHANNEL",
			"ERR_TOOMANYCHANNELS",
			"ERR_TOOMANYTARGETS",
			"ERR_UNAVAILRESOURCE",
			"RPL_TOPIC",
		},
	},

	{
		"PART",
		"<channel> *( \",\" <channel> ) [ <Part Message> ]",
		[]string{
			"ERR_NEEDMOREPARAMS",
			"ERR_NOSUCHCHANNEL",
			"ERR_NOTONCHANNEL",
		},
	},

	{
		"TOPIC",
		"<channel> [ <topic> ]",
		[]string{
			"ERR_NEEDMOREPARAMS",
			"ERR_NOTONCHANNEL",
			"RPL_NOTOPIC",
			"RPL_TOPIC",
			"ERR_CHANOPRIVSNEEDED",
			"ERR_NOCHANMODES",
		},
	},

	{
		"NAMES",
		"[ <channel> *( \",\" <channel> ) [ <target> ] ]",
		[]string{
			"ERR_TOOMANYMATCHES",
			"ERR_NOSUCHSERVER",
			"RPL_NAMREPLY",
			"RPL_ENDOFNAMES",
		},
	},

	{
		"LIST",
		"[ <channel> *( \",\" <channel> ) [ <target> ] ]",
		[]string{
			"ERR_TOOMANYMATCHES",
			"ERR_NOSUCHSERVER",
			"RPL_LIST",
			"RPL_LISTEND",
		},
	},

	{
		"INVITE",
		"<nickname> <channel>",
		[]string{
			"ERR_NEEDMOREPARAMS",
			"ERR_NOSUCHNICK",
			"ERR_NOTONCHANNEL",
			"ERR_USERONCHANNEL",
			"ERR_CHANOPRIVSNEEDED",
			"RPL_INVITING",
			"RPL_AWAY",
		},
	},

	{
		"KICK",
		"<channel> *( \",\" <channel> ) <user> *( \",\" <user> ) [<comment>]",
		[]string{
			"ERR_NEEDMOREPARAMS",
			"ERR_NOSUCHCHANNEL",
			"ERR_BADCHANMASK",
			"ERR_CHANOPRIVSNEEDED",
			"ERR_USERNOTINCHANNEL",
			"ERR_NOTONCHANNEL",
		},
	},

	{
		"PRIVMSG",
		"<msgtarget> <text to be sent>",
		[]string{
			"ERR_NORECIPIENT",
			"ERR_NOTEXTTOSEND",
			"ERR_CANNOTSENDTOCHAN",
			"ERR_NOTOPLEVEL",
			"ERR_WILDTOPLEVEL",
			"ERR_TOOMANYTARGETS",
			"ERR_NOSUCHNICK",
			"RPL_AWAY",
		},
	},

	{
		"NOTICE",
		"<msgtarget> <text>",
		[]string{
			"RPL_MOTDSTART",
			"RPL_MOTD",
			"RPL_ENDOFMOTD",
			"ERR_NOMOTD",
		},
	},

	{
		"LUSERS",
		"[ <mask> [ <target> ] ]",
		[]string{
			"RPL_LUSERCLIENT",
			"RPL_LUSEROP",
			"RPL_LUSERUNKOWN",
			"RPL_LUSERCHANNELS",
			"RPL_LUSERME",
			"ERR_NOSUCHSERVER",
		},
	},

	{
		"VERSION",
		"[ <target> ]",
		[]string{
			"ERR_NOSUCHSERVER",
			"RPL_VERSION",
		},
	},

	{
		"STATS",
		"[ <query> [ <target> ] ]",
		[]string{
			"ERR_NOSUCHSERVER",
			"RPL_STATSLINKINFO",
			"RPL_STATSUPTIME",
			"RPL_STATSCOMMANDS",
			"RPL_STATSOLINE",
			"RPL_ENDOFSTATS",
		},
	},

	{
		"LINKS",
		"[ [ <remote server> ] <server mask> ]",
		[]string{
			"ERR_NOSUCHSERVER",
			"RPL_LINKS",
			"RPL_ENDOFLINKS",
		},
	},

	{
		"TIME",
		"[ <target> ]",
		[]string{
			"ERR_NOSUCHSERVER",
			"RPL_TIME",
		},
	},

	{
		"CONNECT",
		"<target server> <port> [ <remote server> ]",
		[]string{
			"ERR_NOSUCHSERVER",
			"ERR_NOPRIVILEGES",
			"ERR_NEEDMOREPARAMS",
		},
	},

	{
		"TRACE",
		"[ <target> ]",
		[]string{
			"ERR_NOSUCHSERVER",
		},
	},

	{
		"ADMIN",
		"[ <target> ]",
		[]string{
			"ERR_NOSUCHSERVER",
			"RPL_ADMINME",
			"RPL_ADMINLOC1",
			"RPL_ADMINLOC2",
			"RPL_ADMINEMAIL",
		},
	},

	{
		"INFO",
		"[ <target> ]",
		[]string{
			"ERR_NOSUCHSERVER",
			"RPL_INFO",
			"RPL_ENDOFINFO",
		},
	},

	{
		"SERVLIST",
		"[ <mask> [ <type> ] ]",
		[]string{
			"RPL_SERVLIST",
			"RPL_SERVLISTEND",
		},
	},

	{
		"SQUERY",
		"<servicename> <text>",
		[]string{
			"ERR_NORECIPIENT",
			"ERR_NOTEXTTOSEND",
			"ERR_CANNOTSENDTOCHAN",
			"ERR_NOTOPLEVEL",
			"ERR_WILDTOPLEVEL",
			"ERR_TOOMANYTARGETS",
			"ERR_NOSUCHNICK",
			"RPL_AWAY",
		},
	},

	{
		"WHO",
		"[ <mask> [ \"o\" ] ]",
		[]string{
			"ERR_NOSUCHSERVER",
			"RPL_WHOREPLY",
			"RPL_ENDOFWHO",
		},
	},

	{
		"WHOIS",
		"[ <target> ] <mask> *( \",\" <mask> )",
		[]string{
			"ERR_NOSUCHSERVER",
			"ERR_NONICKNAMEGIVEN",
			"RPL_WHOISUSER",
			"RPL_WHOISCHANNELS",
			"RPL_WHOISCHANNELS",
			"RPL_WHOISSERVER",
			"RPL_AWAY",
			"RPL_WHOISOPERATOR",
			"RPL_WHOISIDLE",
			"ERR_NOSUCHNICK",
			"RPL_ENDOFWHOIS",
		},
	},

	{
		"WHOWAS",
		"<nickname> *( \",\" <nickname> ) [ <count> [ <target> ] ]",
		[]string{
			"ERR_NONICKNAMEGIVEN",
			"ERR_WASNOSUCHNICK",
			"RPL_WHOWASUSER",
			"RPL_WHOISSERVER",
			"RPL_ENDOFWHOWAS",
		},
	},

	{
		"KILL",
		"<nickname> <comment>",
		[]string{
			"ERR_NOPRIVILEGES",
			"ERR_NEEDMOREPARAMS",
			"ERR_NOSUCHNICK",
			"ERR_CANTKILLSERVER",
		},
	},

	{
		"PING",
		"<server1> [ <server2> ]",
		[]string{
			"ERR_NOORIGIN",
			"ERR_NOSUCHSERVER",
		},
	},

	{
		"PONG",
		"<server> [ <server2> ]",
		[]string{
			"ERR_NOORIGIN",
			"ERR_NOSUCHSERVER",
		},
	},

	{
		"ERROR",
		"<error message>",
		nil,
	},

	{
		"AWAY",
		"[ <text> ]",
		[]string{
			"RPL_UNAWAY",
			"RPL_NOWAWAY",
		},
	},

	{
		"REHASH",
		"",
		[]string{
			"RPL_REHASHING",
			"ERR_NOPRIVILEGES",
		},
	},

	{
		"DIE",
		"",
		[]string{
			"ERR_NOPRIVILEGES",
		},
	},

	{
		"RESTART",
		"",
		[]string{
			"ERR_NOPRIVILEGES",
		},
	},

	{
		"SUMMON",
		"<user> [ <target> [ <channel> ] ]",
		[]string{
			"ERR_NORECIPIENT",
			"ERR_FILEERROR",
			"ERR_NOLOGIN",
			"ERR_NOSUCHSERVER",
			"ERR_SUMMONDISABLED",
			"RPL_SUMMONING",
		},
	},

	{
		"USERS",
		"[ <target> ]",
		[]string{
			"ERR_NOSUCHSERVER",
			"ERR_FILEERROR",
			"RPL_USERSSTART",
			"RPL_USERS",
			"RPL_NOUSERS",
			"RPL_ENDOFUSERS",
			"ERR_USERSDISABLED",
		},
	},

	{
		"WALLOPS",
		"<Text to be sent>",
		[]string{
			"ERR_NEEDMOREPARAMS",
		},
	},

	{
		"USERHOST",
		"<nickname> *( SPACE <nickname> )",
		[]string{
			"RPL_USERHOST",
			"ERR_NEEDMOREPARAMS",
		},
	},

	{
		"ISON",
		"<nickname> *( SPACE <nickname> )",
		[]string{
			"RPL_ISON",
			"ERR_NEEDMOREPARAMS",
		},
	},
}
