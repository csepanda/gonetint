echo "Start system test"

is_server_running() { [ `ps aux | grep net_server | wc -l` -eq 2 ]; }

panic() { echo "$*" >&2; exit 42; exit 2; }

./net_server &

is_server_running || panic fail: server isn\'t running

echo "[test 1]: client without arguments"
./net_client; EC=$?; [ $EC -eq 2 ] && echo OK || echo "Fail: expected exit code 2 got $EC"
echo "[test 3]: client withot command"
./net_client -port 9000; EC=$?; [ $EC -eq 3 ] && echo OK || echo "Fail: expected exit code 3 got $EC"
echo "[test 4]: client with wrong command usage"
./net_client show; EC=$?; [ $EC -eq 4 ] && echo OK || echo "Fail: expected exit code 4 got $EC"
echo "[test 5]: client list"
./net_client list >&2 && echo OK || echo "Fail"
echo "[test 6]: client show lo"
./net_client show lo >&2 && echo OK || echo "Fail"

pkill net_server
