param([Int32]$p=4)
#Run with -p <number of clients>

$names = @('Gustav', 'Ahmed', 'Simon', 'Alyson', 'Julie', 'Lauritz', 'Marcus', 'Nicklas')
$port = '8080'

$Command = 'cmd /c start powershell -NoExit -Command {
    $host.UI.RawUI.WindowTitle = "Server - Chitty-Chat";
    $host.UI.RawUI.BackgroundColor = "black";
    $host.UI.RawUI.ForegroundColor = "white";
    Clear-Host;
    cd server; 
    go run . -port ' + $port + ';
}'

invoke-expression -Command $Command;

for ($i = 0; $i -lt $p; $i++) {
    $name = $i
    
    if ($i -lt $names.count) {
        $name = $names[$i]
    }

    $Command = 'cmd /c start powershell -NoExit -Command {
        $host.UI.RawUI.WindowTitle = "Client - ' + $name + '";
        cd client; 
        go run . -name ' + $name +' -server ' + $port +';
    }'
    
    invoke-expression -Command $Command
};