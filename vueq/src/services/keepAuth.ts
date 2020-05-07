
import { instance as store } from 'src/store'

async function sleepAndTrigger (seconds: number, callback: Function) {
  await new Promise((resolve) => setTimeout(resolve, seconds * 1000))
  callback()
}
export class KeepAuth {
  private static instance: KeepAuth

  public static readonly TimeSeconds = 60 * 60

  public static fire () {
    if (!KeepAuth.instance) {
      KeepAuth.instance = new KeepAuth()
    }
    return KeepAuth.instance
  }

  private constructor() {
    this.mainLoop()
  }

  private mainLoop () {
    this.checkAuth()
    sleepAndTrigger(KeepAuth.TimeSeconds, () => {
      this.mainLoop()
    })
  }

  private checkAuth () {
    store.dispatch('user/REFRESH')
  }
}
