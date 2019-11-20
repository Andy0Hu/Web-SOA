package com.soaproject.demo.server.request;

import lombok.Data;

import javax.validation.constraints.NotNull;

@Data
public class CheckRequest {
    @NotNull
    private String orderId;
}
