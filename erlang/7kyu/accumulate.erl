-module(accumulate).
-export([accumulate/2]).

-spec accumulate(fun((A) -> B), list(A)) -> list(B).

% accumulate(_Fn, [], Acc) -> lists:reverse(Acc);
% accumulate(Fn, [Head|Tail], Acc) -> accumulate(Fn, Tail, [Fn(Head) | Acc]).
% accumulate(Fn, List) -> accumulate(Fn, List, []).

% accumulate(_Fn, []) -> [];
% accumulate(Fn, [Head|Tail]) ->
%     [Fn(Head) | accumulate(Fn, Tail)].

accumulate(Fn, List) -> [Fn(Char) || Char <- List].
