New  env setup


add to settings.json
{
    "go.goroot": "c:/sdk/go"
}
and restart vscode


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
