param([Int32]$r=3, [Int32]$n=2)
# r = Number of Replicas
# n = Number of Client nodes to connect

$replicaPorts = @(8080, 8081, 8082, 8083, 8084)
$frontEndPorts = @(8085, 8086, 8087, 8088, 8089)
$command = ''

"STARTING GO"

# Start replicas
for ($i = 0; $i -lt $r; $i++) {
    if ($i -gt $replicaPorts.count) {
        "WARNING: Max number of replicas for this script has been created (" + $replicaPorts.count + "). => Not creating any more replicas!"
        break
    }

    $port = $replicaPorts[$i]

    "REPLICA => Server port: $port |"

    $command = 'cmd /c start powershell -NoExit -Command {
                $host.UI.RawUI.WindowTitle = "Node - ' + $name + '";
                $host.UI.RawUI.BackgroundColor = "black";
                $host.UI.RawUI.ForegroundColor = "white";
                Clear-Host;
                cd node;
                go run . -name ' + $name + ' -sport ' + $sport + ' -ips ' + $ips + ' -delay ' + $delay + ';
            }'

    Invoke-Expression -Command $command
}

# Start Clients
for ($i = 0; $i -lt $n; $i++) {
    if ($i -gt $frontEndPorts.count) {
        "WARNING: Max number of clients for this script has been created (" + $frontEndPorts.count + "). => Not creating any more clients!"
        break
    }

    $port = $frontEndPorts[$i]

    "CLIENT => Server port: $port |"

    $command = 'cmd /c start powershell -NoExit -Command {
                $host.UI.RawUI.WindowTitle = "Node - ' + $name + '";
                $host.UI.RawUI.BackgroundColor = "black";
                $host.UI.RawUI.ForegroundColor = "white";
                Clear-Host;
                cd node;
                go run . -name ' + $name + ' -sport ' + $sport + ' -ips ' + $ips + ' -delay ' + $delay + ';
            }'

    Invoke-Expression -Command $command
}