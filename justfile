echo:
  cd echo && go build
  ~/Downloads/maelstrom/maelstrom test -w echo --bin ./echo/echo --node-count 1 --time-limit 10

uniqueids:
  cd uniqueids && go build
  ./maelstrom test -w unique-ids --bin ./uniqueids/uniqueids --time-limit 30 --rate 1000 --node-count 3 --availability total --nemesis partition

debug:
  ~/Downloads/maelstrom/maelstrom serve