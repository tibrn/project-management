import { instance as axios } from './../boot/axios'
import { Notify, Dialog } from 'quasar'

interface CallbackFn<T> {
  (...args: any[]): T;
}

type CallbackString = CallbackFn<string>

interface OptionalOptions {
  afterDelete?: CallbackFn<any>;
  beforeDelete?: CallbackFn<any>;
  title?: string | CallbackString;
  message?: string | CallbackString;
}

export interface OptionsDelete extends OptionalOptions {
  deleteRoute: string | CallbackString;
}

function defaultsOrOptions (def: OptionalOptions, options: OptionsDelete): OptionsDelete {
  if (!options.afterDelete) options.afterDelete = def.afterDelete
  if (!options.title) options.title = def.title
  if (!options.message) options.message = def.message
  return options
}

const notifyOptions = Object.freeze({
  message: 'Delete canceled',
  timeout: 2500,
  progress: true,
  position: 'top',
  color: 'secondary',
})

const cancelNotification = () => {
  // Cancel notification
  Notify.create(notifyOptions)
}

const defaultOptions = Object.freeze({
  title: 'Delete',
  message: 'Are you sure that you want to delete ?',
})

export function createDelete (options: OptionsDelete) {
  // take defaults
  options = defaultsOrOptions(defaultOptions, options)

  const deleteCallback = async (...args: unknown[]) => {
    try {
      // BeforeDelete
      if (typeof options.beforeDelete === 'function') {
        await options.beforeDelete(...args)
      }

      // Delete entity from server
      await axios.delete(typeof options.deleteRoute === 'string' ? options.deleteRoute : options.deleteRoute(...args))

      // AfterDelete with success
      if (typeof options.afterDelete === 'function') {
        await options.afterDelete(...args)
      }
    } catch (e) { }
  }

  const confirmDelete = (...args: unknown[]) => {
    // Confirm dialog
    Dialog.create({
      title: typeof options.title !== 'function' ? options.title : options.title(...args),
      message: typeof options.message !== 'function' ? options.message : options.message(...args),
      ok: 'Delete',
      cancel: 'Cancel',
    })
      .onOk(() => deleteCallback(...args))
      .onCancel(cancelNotification)
  }

  return {
    confirmDelete,
  }
}
