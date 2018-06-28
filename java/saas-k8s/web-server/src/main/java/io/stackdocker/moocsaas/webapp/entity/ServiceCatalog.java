package io.stackdocker.moocsaas.webapp.entity;

import java.util.ArrayList;
import java.util.Calendar;
import java.util.List;

import javax.validation.constraints.NotEmpty;

public class ServiceCatalog {

        private Long id;

        private String name;

        @NotEmpty(message = "Text is required.")
        private String text;

        @NotEmpty(message = "Summary is required.")
        private String summary;

        private Calendar created = Calendar.getInstance();

        public Long getId() {
            return this.id;
        }

        public void setId(Long id) {
            this.id = id;
        }

        public Calendar getCreated() {
            return this.created;
        }

        public void setCreated(Calendar created) {
            this.created = created;
        }

        public String getText() {
            return this.text;
        }

        public void setText(String text) {
            this.text = text;
        }

        public String getSummary() {
            return this.summary;
        }

        public void setSummary(String summary) {
            this.summary = summary;
        }

        public String getName() {
            return this.name;
        }

        public void setName(String name) {
            this.name = name;
        }

        private List<ServiceInstance> instances = new ArrayList<>();

}
