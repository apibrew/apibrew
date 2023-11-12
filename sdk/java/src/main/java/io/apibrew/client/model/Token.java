package io.apibrew.client.model;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.Date;

public class Token {

    public static class Container {
        private Token token;

        public Token getToken() {
            return token;
        }

        public void setToken(Token token) {
            this.token = token;
        }
    }

    public enum TokenTerm {
        // 1 minute
        VERY_SHORT,
        // 20 minutes
        SHORT,
        // 2 days
        MIDDLE,
        // 60 days
        LONG,
        // 2 years
        VERY_LONG,
    }

    @JsonProperty
    private TokenTerm term;

    @JsonProperty
    private String content;

    @JsonProperty
    private Date expiration;
    public Date getExpiration() {
        return expiration;
    }

    public void setExpiration(Date expiration) {
        this.expiration = expiration;
    }

    public String getContent() {
        return content;
    }

    public void setContent(String content) {
        this.content = content;
    }

    public TokenTerm getTerm() {
        return term;
    }

    public void setTerm(TokenTerm term) {
        this.term = term;
    }
}
