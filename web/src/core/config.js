/**
 * 网站配置文件
 */

const config = {
  appName: '博远天合OA',
  appLogo: 'https://byapp.bythedu.com/media/logo.png',
  showViteLogo: true
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    const chalk = require('chalk')
    console.log(
      chalk.green(
        `欢迎使用ByOA`
      )
    )
    
  }
}

export default config
