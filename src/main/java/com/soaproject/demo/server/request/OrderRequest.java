package com.soaproject.demo.server.request;

import lombok.Data;
import javax.validation.constraints.NotNull;
import java.math.BigDecimal;

@Data
public class OrderRequest  {
    @NotNull
    private String orderId;
    @NotNull
    private String userId;
    @NotNull
    private String destination;
    @NotNull
    private String orderDescription;
    @NotNull
    private BigDecimal weight;




}
