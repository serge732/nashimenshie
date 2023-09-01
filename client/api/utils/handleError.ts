import { AxiosError } from 'axios';

import { RequestResult } from '../types';

export const handleError = (error?: unknown): RequestResult<null> => {
    return {
        ok: false,
        data: null,
        error: (() => {
            if (error instanceof AxiosError) {
                if (error.response?.status === 500) {
                    return 'Не удалось получить ответ от сервера';
                }
            }

            if (error instanceof Error) {
                return error.message;
            }

            return 'Произошла непредвиденная ошибка';
        })(),
    };
};
