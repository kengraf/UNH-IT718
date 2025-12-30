# IT718 – Deploy a Static Web Site
## Objectives
For each of the three major cloud vendors (AWS, Azure, and GCP) complete the following actions:
1.	Identify static web site you would like to publish on the web.
2.	Push the web content to cloud storage.
3.	Make the cloud storage publicly available.

> [!NOTE]
> Running the standalone dedicated nginx server you created in the last lab is generally a bad idea for serving a static web site.

Standalone servers have maintenance, monitoring, and run-time costs that create unnecessary overhead for a static public website.  

Even if your website does more than simple GETs, offloading the static content (html, css, javascript, images) to cloud based CDN services will offer better performance, availability, and reduced cost.  

A simple "Hello World" webpage is enough to complete this lab.  But do your lab grade, yourself, and your semester end project a favor by pushing something more substantial.  Ideally, this is a good time to start building out your project by understanding how to publish web content.  

If you don't have existing content, a Google search for templates will find many options/  I have had success with these sources in the past: *[Github](https://github.com/website-templates)* and *[Html5Up](https://html5up.net/)*.  


__Highly recommended but optional:__
Use a custom DNS domain (friendly URL) for your site.  All 3 vendors charge $0.50/month to host a DNS zone.  AWS and GCP can act as your DNS registrar for ~$12.year.  Yes, I know sites like <your-name>.github.io are free.  But, being able to talk and show that you understand the entire CI/CD process are great talking points during a hiring interview.  


> [!Tip]
> In the later Github lab you will automatically push content from a github repo eliminating the need for using a cloud shell to push content by hand.  
## Lab Report
For each cloud vendor show:
1.	The output from the cloud shell showing a wget command retrieving the default page.
2.	Screenshot of browser rendering your home page.
3.	The URL to retrieve your website. Provided as text, not a screenshot, in your lab report
   * The website needs to remain up until the lab is graded.
