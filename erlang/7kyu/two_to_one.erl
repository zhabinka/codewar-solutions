-module(two_to_one).
-export([longest/2]).

-spec longest(string, string) -> string.

longest(L, M) ->
	% First solution
	% UList = lists:uniq([[Char] || Char <- L ++ M]),
	% string:join(lists:sort(UList), "").

	lists:usort(L ++ M).
