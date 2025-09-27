-module(easyline).
-export([easyline/1]).

-spec(easyline(integer()) -> integer()).

% Сумма квадрато биномиальных коэффициентов 
% (2n)!/(n!)2

fact(0) ->
  1;

fact(Num) when Num > 0, is_integer(Num) ->
  Num * fact(Num - 1).

easyline(Num) ->
  fact(2 * Num) div (fact(Num) * fact(Num)).
