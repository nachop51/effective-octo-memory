export async function handlePromise<T>(promise: Promise<T>): Promise<
	| {
			ok: true
			data: T
	  }
	| { ok: false; error: Error }
> {
	try {
		const result = await promise

		return {
			ok: true,
			data: result
		}
	} catch (error) {
		return {
			ok: false,
			error: error instanceof Error ? error : new Error('Unknown error')
		}
	}
}
