export class NotFoundError extends Error {
  constructor(message: string = 'Not Found') {
    super(message)
    this.name = 'NotFoundError'
  }
}

export class BadRequestError extends Error {
  constructor(message: string = 'Bad Request') {
    super(message)
    this.name = 'BadRequestError'
  }
}

export class InternalServerError extends Error {
  constructor(message: string = 'Internal Server Error') {
    super(message)
    this.name = 'InternalServerError'
  }
}
