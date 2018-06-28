/*
  Inspired by:
  - https://github.com/spring-projects/spring-boot/blob/master/spring-boot-samples/spring-boot-sample-web-ui/src/main/java/sample/web/ui/mvc/MessageController.java
*/
package io.stackdocker.moocsaas.webapp.mvc;

import javax.validation.Valid;

import com.google.common.collect.ImmutableMap;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.validation.BindingResult;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.servlet.ModelAndView;
import org.springframework.web.servlet.mvc.support.RedirectAttributes;

import io.stackdocker.moocsaas.webapp.entity.*;
import io.stackdocker.moocsaas.webapp.service.*;

@Controller
@RequestMapping("/")
public class SampleController {
    private static final Logger logger = LoggerFactory.getLogger(SampleController.class);

	private final CatalogRepository catalogRepository;
	private final InstanceRepository instanceRepository;

	@Autowired
	public SampleController(CatalogRepository catalogRepository, InstanceRepository instanceRepository) {
		this.catalogRepository = catalogRepository;
		this.instanceRepository = instanceRepository;

        ServiceCatalog sc = new ServiceCatalog();
        sc.setName("ctf-v3");
        sc.setText("CTF-Java");
        sc.setSummary("CTF-Java");
        sc = this.catalogRepository.save(sc);
	}

	@GetMapping("/")
	public ModelAndView list() {
		Iterable<ServiceCatalog> scs = this.catalogRepository.findAll();
        ServiceCatalog sc = scs.iterator().next();
		return new ModelAndView("messages/list", ImmutableMap.of("sc", sc, "scs", scs));
	}

	@GetMapping("{id}")
	public ModelAndView view(@PathVariable("id") ServiceCatalog sc) {
		return new ModelAndView("messages/view", "message", sc);
	}


    @GetMapping(value = "/", params = "form")
    public String createForm(@ModelAttribute ServiceCatalog sc, Model model) {
	    logger.debug("form", sc);
	    ServiceInstance si = new ServiceInstance();
	    si.setName("demo");
	    si.setText("a CTF demo");
	    si.setSummary("a CTF-Java docker cluster");
	    si.setCatalogId(this.catalogRepository.findAll().iterator().next().getId());
	    model.addAttribute("si", si);
        return "messages/form";
    }

    @PostMapping("/")
	public ModelAndView create(@Valid ServiceInstance si, BindingResult result,
			RedirectAttributes redirect) {
		if (result.hasErrors()) {
			return new ModelAndView("messages/form", "formErrors", result.getAllErrors());
		}
        si = this.instanceRepository.saveIntoCreation(si);
		redirect.addFlashAttribute("globalMessage", "Successfully created a new message");
        ServiceCatalog sc = this.catalogRepository.findAll().iterator().next();
		return new ModelAndView("redirect:/{sc.id}", "sc.id", sc.getId());
	}

	@RequestMapping("foo")
	public String foo() {
		throw new RuntimeException("Expected exception in controller");
	}

	@GetMapping("delete/{id}")
	public ModelAndView delete(@PathVariable("id") Long id) {
		this.instanceRepository.deleteServiceInstance(id);
		Iterable<ServiceInstance> messages = this.instanceRepository.findAll();
		return new ModelAndView("messages/list", "messages", messages);
	}

	@GetMapping("modify/{id}")
	public ModelAndView modifyForm(@PathVariable("id") ServiceCatalog message) {
		return new ModelAndView("messages/form", "message", message);
	}

}