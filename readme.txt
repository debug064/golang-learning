New  env setup

1. Install MSYS2
2. Run MSYS2 XX64
3. 
 $ pacman -Syu
 $ pacman -S git mingw-w64-x86_64-toolchain mingw-w64-x86_64-go

4. Check go env
5. Modify .bachrc

$echo "export PATH=\$PATH:/D:/Programs/Microsoft VS Code/bin" >> ~/.bashrc

export GOPATH=~/go
export PATH=$PATH:~/go/bin
export PATH=$PATH:"/d/Programs/Microsoft VS Code/bin"

6. modify go env
go env -w GOCACHE=~/go-build
and other if they point outside MSYS2


7. Check
echo $PATH

8. git
  git config --global user.email "you@example.com"
  git config --global user.name "Your Name"

to set your account's default identity.
Omit --global to set the identity only in this repository.

git config user.name "sam"
git config user.email "sam@sam"
git config --list

https://stackoverflow.com/questions/71814348/how-to-run-go-file-on-vscode

https://stackoverflow.com/questions/66894200/error-message-go-go-mod-file-not-found-in-current-directory-or-any-parent-dire/67598174#67598174
go env -w GO111MODULE=off


add to settings.json
{
    "go.goroot": "c:/sdk/go"
}
and restart vscode



for c++
pacman -S mingw-w64-x86_64-toolchain