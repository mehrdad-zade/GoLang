Book Store app API:
- MySQL DB
- CRUD APIs: fid the possible routs under resources/Routs.png.
- GORM package to interact with the DB
- Json marshall, unmarshall
- Project structure is done properly rather than having everything in a main function. Please find the structure under resources/ProjectStructure.png
- Gorilla Mux

Packages:
- //go mod init in the root project.
- go get "github.com/jinzhu/gorm"
- go get "github.com/jinzhu/gorm/dialects/mysql"
- go get "github.com/gorilla/mux"


Note:
if you have an error like "gopls was not able to find modules in your workspace":
- go to, File > Preferences > Settings
- search for "gopls"
- Edit in settings.json
- "gopls": {
    "experimentalWorkspaceModule": true,
}
- restart VCS
