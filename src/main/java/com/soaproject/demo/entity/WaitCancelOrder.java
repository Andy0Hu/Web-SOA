package com.soaproject.demo.entity;

public class WaitCancelOrder {
    /**
     * This field was generated by MyBatis Generator.
     * This field corresponds to the database column waitcancelorder.task_id
     *
     * @mbggenerated
     */
    private String ctaskId;

    /**
     * This field was generated by MyBatis Generator.
     * This field corresponds to the database column waitcancelorder.corder_id
     *
     * @mbggenerated
     */
    private String corderId;

    /**
     * This field was generated by MyBatis Generator.
     * This field corresponds to the database column waitcancelorder.cuser_id
     *
     * @mbggenerated
     */
    private String cuserId;
    private String taskId;

    /**
     * This method was generated by MyBatis Generator.
     * This method returns the value of the database column waitcancelorder.ctask_id
     *
     * @return the value of waitcancelorder.ctask_id
     *
     * @mbggenerated
     */
    public String getCtaskId() {
        return ctaskId;
    }

    /**
     * This method was generated by MyBatis Generator.
     * This method sets the value of the database column waitcancelorder.task_id
     *
     * @param ctaskId the value for waitcancelorder.task_id
     *
     * @mbggenerated
     */
    public void setCtaskId(String ctaskId) {
        this.ctaskId = ctaskId == null ? null : ctaskId.trim();
    }

    /**
     * This method was generated by MyBatis Generator.
     * This method returns the value of the database column waitcancelorder.corder_id
     *
     * @return the value of waitcancelorder.corder_id
     *
     * @mbggenerated
     */
    public String getCorderId() {
        return corderId;
    }

    /**
     * This method was generated by MyBatis Generator.
     * This method sets the value of the database column waitcancelorder.corder_id
     *
     * @param corderId the value for waitcancelorder.corder_id
     *
     * @mbggenerated
     */
    public void setCorderId(String corderId) {
        this.corderId = corderId == null ? null : corderId.trim();
    }

    /**
     * This method was generated by MyBatis Generator.
     * This method returns the value of the database column waitcancelorder.cuser_id
     *
     * @return the value of waitcancelorder.cuser_id
     *
     * @mbggenerated
     */
    public String getCuserId() {
        return cuserId;
    }

    /**
     * This method was generated by MyBatis Generator.
     * This method sets the value of the database column waitcancelorder.cuser_id
     *
     * @param cuserId the value for waitcancelorder.cuser_id
     *
     * @mbggenerated
     */
    public void setCuserId(String cuserId) {
        this.cuserId = cuserId == null ? null : cuserId.trim();
    }

    public String getTaskId() {
        return taskId;
    }
    public void setTaskId(String taskId) {
        this.taskId = taskId == null ? null : taskId.trim();
    }
}