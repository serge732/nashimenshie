export type RequestResult<T> = {
    ok: boolean;
    data: T | null;
    error: string | null;
};
