default:
	rm -rf /tmp/fossora/default.etcd
	go run cmd/fossora/main.go metta

pi0:
	GOARCH=arm64 GOOS=linux go build -o build/linux/arm64/fossora cmd/fossora/main.go
	ssh sawyer@pi0 "rm -rf /home/sawyer/repos; mkdir -p repos/fossora; rm -rf /tmp/fossora/default.etcd;"
	scp build/linux/arm64/fossora sawyer@pi0:/home/sawyer/repos/fossora/
	ssh sawyer@pi0 "chmod +x /home/sawyer/repos/fossora/fossora"
	ssh -t sawyer@pi0 "/home/sawyer/repos/fossora/fossora pi0"

pi1:
	GOARCH=arm64 GOOS=linux go build -o build/linux/arm64/fossora cmd/fossora/main.go
	ssh sawyer@pi1 "rm -rf /home/sawyer/repos; mkdir -p repos/fossora; rm -rf /tmp/fossora/default.etcd;"
	scp build/linux/arm64/fossora sawyer@pi1:/home/sawyer/repos/fossora/
	ssh sawyer@pi1 "chmod +x /home/sawyer/repos/fossora/fossora"
	ssh -t sawyer@pi1 "/home/sawyer/repos/fossora/fossora pi1"
