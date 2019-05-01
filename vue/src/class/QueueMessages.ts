//QueueMessages is used to show gloabl messages
export class QueueMessages {
  private static instance: QueueMessages;
  private messages: Array<Message>;
  private isMessage: boolean;
  private timeout: number;
  private queueTimeout: number;
  public message?: Message;
  private constructor(timeout: number) {
    this.messages = [];
    this.isMessage = false;
    this.timeout = timeout;
    this.queueTimeout = -1;
  }

  //Return instace of queue
  public static getInstance(timeout?: number) {
    if (QueueMessages.instance === undefined) {
      if (typeof timeout === "number")
        QueueMessages.instance = new QueueMessages(timeout);
      console.log("QUEUE MESSAGE");
      console.log(QueueMessages.instance);
    }
    return QueueMessages.instance;
  }

  //Send message to queue
  //Triqqers showMessage if no message is shown
  public sendMessage(message: Message) {
    this.messages.push(message);
    if (!this.isMessage) this.showMessage();
  }

  //Show message from queue
  //Recursiv implementation
  private showMessage() {
    console.log(this);
    if (this.messages.length > 0) {
      this.message = this.messages[0];
      this.isMessage = true;
      this.messages.splice(0, 1);
      let self = this;
      this.queueTimeout = setTimeout(function() {
        self.showMessage();
      }, this.timeout);
    } else if (this.isMessage) {
      this.message = undefined;
      this.isMessage = false;
    }
  }

  //Close current message
  //and go to next if exist
  public closeMessage() {
    clearTimeout(this.queueTimeout);
    this.showMessage();
  }
}

export interface Message {
  "message-type": string;
  text: string;
}
