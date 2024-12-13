echo:
  cd echo && go build
  ~/Downloads/maelstrom/maelstrom test -w echo --bin ./echo/echo --node-count 1 --time-limit 10

debug:
  ~/Downloads/maelstrom/maelstrom serve