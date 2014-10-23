Move fast and don't break your infra configuraion
====

- Ok, Move fast and don’t break your infra configuration
- Hi, This is me
  - I'm known as deeeet in twitter
  - And I’m working as a Software engineer in Rakuten
- This is my Japanese blog
  - I write something about Docker or Golang
  - If you like my presentation today, Please check it
- And I build some tool
  - This is cli-init
    - It provides the easy way to start building Golang command-line application
  - And this is GHR
    - It enables you to release your tool to Github very easily
- Ok, I’m gonna Talk About Docker
  - But Today is comunities meetup. So I would like to talk tool related to Chef or Puppet or OpenStack
- So do you think Container is future ?
  - Google says:
    - Every thing at Google runs in a container
    - And They deploy over 2billion containers per week
    - https://speakerdeck.com/jbeda/containers-at-scale (https://speakerdeck.com/jbeda/containers-at-scale)
  - And recently Microsoft annouced:
    - Integrate Docker with Windows Server 		- http://weblogs.asp.net/scottgu/docker-and-microsoft-integrating-docker-with-windows-server-and-microsoft-azure (http://weblogs.asp.net/scottgu/docker-and-microsoft-integrating-docker-with-windows-server-and-microsoft-azure)
- It’s really hot, so we should try it
  - Container deployment make our business great or not
- But How ?
  - There are some risks and We may already have its own infrastructure configuration like Chef, puppet or Ansible
- One of my opinion is:
  - using Packer
- Packer is:
  - an open source tool for creating machine images from a single json configuration file.
- Packer creates image for
  - Docker and OpenStack, Amazon AMI and so on
- And packer is able to use Chef or Puppet or Ansible for machine provisioning
- It means you can build docker without Dockerfile but with your configuration tool
- It’s very important because If you don't like Docker, Dockerfiles have to be translated over to another format
- You can easily pivot Without Breaking Your Infra Configuration
- Ok, Let’s build docker image with your chef cookbook
- All we need to do is to prepare a single json file and our chef cookbook
- In json file, we need to write 3 parts, builders and provisioner and post-processor
- In builder section, write base image to start with
  - In this case, I'm providing ubuntu
- In provisioner section, write how to install software and configure it
  - In this case, I'm provide chef cookbook path and run-list for it
- In post-processor section, write how to generate machine image
  - In this case, I'm commiting with tag 
- That’s it every simple
- And after that, type packer build machine.json
  - You can get docker image without Dockerfile
- Ok I’ll do some simple domo

```json
 "builders": ()
 "type": "docker",
 "image": "ubuntu:latest",
 "commit": true
 }],

 "provisioners": ()
 
 "type": "shell",
 "inline": "apt-get -y update; apt-get install -y curl" ()
 },
 
 "type": "chef-solo",
 "cookbookpaths": "site-cookbooks" (),
 "runlist": "apachedefault" ()
 }
 ],

 "post-processors":  ()
 
 "type": "docker-import",
 "repository": "tcnksm/packer-chef",
 "tag": "latest"
 },
 "docker-push"
 ]
 }
```



