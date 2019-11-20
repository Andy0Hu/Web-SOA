package com.soaproject.demo.server.listener;

import com.soaproject.demo.dao.OrderMapper;
import com.soaproject.demo.entity.Order;
import org.apache.commons.lang.StringUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.data.redis.core.StringRedisTemplate;
import org.springframework.messaging.handler.annotation.Payload;

import javax.annotation.Resource;
import java.util.Objects;

public class OrderListener {
    private final static Logger log= LoggerFactory.getLogger(OrderListener.class);

    @Resource
    private OrderMapper orderMapper;
    @Resource
    private StringRedisTemplate redisTemplate;

    //直接消费模式
    @RabbitListener(queues = "${register.queue.name}",containerFactory = "singleListenerContainer")
    public void consumeMessage(@Payload Order order){
        try {
            log.debug("消费者监听新建交易记录信息： {} ",order);

            //TODO：表示已经到ttl了，却还没付款，则需要处理为失效
            if (Objects.equals("未完成",order.getOrderStatus())){
                order.setOrderStatus("已完成");
                order.setDeliveryInfo("已送达");
                orderMapper.updateByPrimaryKeySelective(order);
                String s = redisTemplate.opsForValue().get("ORDER");
                if (StringUtils.isNotBlank(s)){
                    redisTemplate.delete("ORDER");
                }
            }
        }catch (Exception e){
            log.error("消息体解析 发生异常； ",e.fillInStackTrace());
        }
    }


}
