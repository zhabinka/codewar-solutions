-module(seven).
-export([seven/1]).

-spec(seven(integer()) -> {integer, integer}).

seven(M) -> seven(M, 0).

seven(Num, Steps) when Num < 100 ->
  {Num, Steps};

seven(Num, Step) ->
  seven(Num div 10 - Num rem 10 * 2, Step + 1).
