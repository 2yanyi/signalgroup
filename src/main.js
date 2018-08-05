const AppUrl = "https://im.dingtalk.com"
    , electron = require('electron'), width = 1000, height = 600
    , main_window = () => {
        electron.Menu.setApplicationMenu(null)
        let mainWindow = new electron.BrowserWindow({
            width: width, height: height, transparent: true, frame: true,
        })
        mainWindow.on('closed', () => { mainWindow = null }).loadURL(AppUrl)
        new_window(mainWindow)
    }
    , new_window = w => {
        w.webContents.on('new-window', (event, url, fname, disposition, options) => {
            event.preventDefault()
            let childWindow = new electron.BrowserWindow({ width: width, height: 800 })
            childWindow.on('closed', () => { mainWindow = null }).loadURL(url)
            new_window(childWindow)
        })
    }
electron.app.on("ready", main_window).on('window-all-closed', () => {
    electron.app.quit()
});
