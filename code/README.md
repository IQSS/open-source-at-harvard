# Code for reproducibility

This "code" directory is a playground for scripts that operate on files in the "data" directory. Like the "data" directory, this code can be (and is being, manually) uploaded to the [Open Source at Harvard dataset][].

[Open Source at Harvard dataset]: https://doi.org/10.7910/DVN/TJCLKP

Ideally, these scripts can be used locally as well as within a computational environment (Binder, Whole Tale, etc.) that has downloaded or otherwise accessed the dataset. Here's a Binder button, for example, that will give you access to the code and data in a terminal:

[![Binder](https://mybinder.org/badge_logo.svg)](https://mybinder.org/v2/dataverse/10.7910/DVN/TJCLKP/)

Because Dataverse changes the .tsv into a .tab ([#2720][]) for now, from within Binder you need to modify the script to use the .tab version (until https://github.com/jupyterhub/repo2docker/pull/1253 is merged). Then you can run it, like this:

```
$ python3 code/language.py 
$ head -4 language.tsv
language        repo
Java    https://github.com/IQSS/dataverse
Python  https://github.com/glue-viz/glue
R       https://github.com/hms-dbmi/UpSetR
```

[#2720]: https://github.com/IQSS/dataverse/issues/2720
