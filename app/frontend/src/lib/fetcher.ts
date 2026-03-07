import { API_BASE_URL } from '$env/static/private';

const fetcher = async (
	method: 'GET' | 'POST' | 'PUT' | 'DELETE',
	route: string,
	token: string,
	init?: RequestInit
) => {
	const fullUrl = `${API_BASE_URL}${route}`;
	const isFormData = init?.body instanceof FormData;
	const headers: HeadersInit = {
		Accept: 'application/json',
		Authorization: `Bearer ${token}`,
		...(isFormData ? {} : { 'Content-Type': 'application/json' })
	};
	const options: RequestInit = {
		method,
		...init,
		headers
	};

	const logResponse = async (response: Response) => {
		let content = 'unknown';
		if (!response.ok) {
			try {
				content = await response.clone().text();
			} catch {
				content = '[Unable to read response body]';
			}
		} else {
			content = 'success';
		}
		const message = {
			timestamp: new Date().toISOString(),
			method,
			url: fullUrl,
			status: response.status,
			body: isFormData ? 'FormData' : (init?.body ?? 'none'),
			response: content
		};
		console.log(JSON.stringify(message));
	};

	try {
		const response = await fetch(fullUrl, options);
		await logResponse(response);
		return response;
	} catch (err) {
		console.error('API fetch error:', err);
		throw err;
	}
};

export const apiGet = async (route: string, token: string, init?: RequestInit) => {
	return fetcher('GET', route, token, init);
};
export const apiPost = async (route: string, token: string, init?: RequestInit) => {
	return fetcher('POST', route, token, init);
};
export const apiPut = async (route: string, token: string, init?: RequestInit) => {
	return fetcher('PUT', route, token, init);
};
export const apiDelete = async (route: string, token: string, init?: RequestInit) => {
	return fetcher('DELETE', route, token, init);
};
