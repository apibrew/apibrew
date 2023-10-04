package io.apibrew.client;

import io.apibrew.client.model.Extension;
import lombok.Getter;

import static io.apibrew.client.model.Extension.Code.INTERNAL_ERROR;

@Getter
public class ApiException extends RuntimeException {

    private final Extension.Error error;

    public ApiException(Extension.Error error) {
        super(error.getMessage());
        this.error = error;
    }

    public ApiException() {
        this.error = new Extension.Error().withCode(INTERNAL_ERROR);
    }

    public ApiException(String message) {
        super(message);
        this.error = new Extension.Error().withCode(INTERNAL_ERROR).withMessage(message);
    }

    public ApiException(Extension.Code code, String message) {
        super(message);
        this.error = new Extension.Error().withCode(code).withMessage(message);
    }

    public ApiException(String message, Throwable cause) {
        super(message, cause);
        this.error = new Extension.Error().withCode(INTERNAL_ERROR).withMessage(message);
    }

    public ApiException(Throwable cause) {
        super(cause);
        this.error = new Extension.Error().withCode(INTERNAL_ERROR).withMessage(cause.getMessage());
    }

    public ApiException(String message, Throwable cause, boolean enableSuppression, boolean writableStackTrace) {
        super(message, cause, enableSuppression, writableStackTrace);
        this.error = new Extension.Error().withCode(INTERNAL_ERROR).withMessage(message);
    }

}
