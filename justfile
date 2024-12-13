MAELSTROM_PATH := (env_var("HOME") + "/Downloads/maelstrom")

echo:
  cd echo && go build
  {{MAELSTROM_PATH}}/maelstrom test -w echo --bin ./echo/echo --node-count 1 --time-limit 10

uniqueids:
  cd uniqueids && go build
  {{MAELSTROM_PATH}}/maelstrom test -w unique-ids --bin ./uniqueids/uniqueids --time-limit 30 --rate 1000 --node-count 3 --availability total --nemesis partition

broadcast:
  cd broadcast && go build
  {{MAELSTROM_PATH}}/maelstrom test -w broadcast --bin ./broadcast/broadcast --node-count 1 --time-limit 20 --rate 10

debug:
  {{MAELSTROM_PATH}}/maelstrom serve