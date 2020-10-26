CurrentVersion=0.2.7
Project=github.com/Kaytz/KayProxy
Path="$Project/version"
ExecName="KayProxy"
GitCommit=$(git rev-parse --short HEAD || echo unsupported)
GoVersion=$(go version)
BuildTime=$(date "+%Y-%m-%d %H:%M:%S")
echo "Building..."
TargetDir=bin
env GOPROXY=https://goproxy.io go build -ldflags "-X '$Path.Version=$CurrentVersion' -X '$Path.BuildTime=$BuildTime' -X '$Path.GoVersion=$GoVersion' -X '$Path.GitCommit=$GitCommit' -w -s" -o $TargetDir/$ExecName
echo "--------------------------------------------"
echo "Version:" $CurrentVersion
echo "Git commit:" "$GitCommit"
echo "Go version:" "$GoVersion"
echo "Build Time:" "$BuildTime"
echo "Build Finish"
echo "--------------------------------------------"
