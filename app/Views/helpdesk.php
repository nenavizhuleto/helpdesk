<?= $this->extend('layouts/main') ?>


<?= $this->section('content') ?>

<div class="card my-4">
    <div class="card-body">
        <h2>Hostname: <?= $hostname ?></h2>
    </div>
</div>

<div class="card my-4">
    <div class="card-body">
        <h2>Address: <?= $address ?></h2>
        <h2>Port: <?= $port ?></h2>
    </div>
</div>


<div class="card my-4">
    <div class="card-body">
        <h2>MAC Address: <?= $mac ?></h2>
    </div>
</div>

<?php foreach ($networks as $network): ?>
<div class="card my-4">
    <div class="card-body">

        <div class="row">
            <div class="col">
                <?= $network['network'] ?>
            </div>
            <div class="col">
                <?= $network['client'] ?>
            </div>
            <div class="col">
                <?= $network['location'] ?>
            </div>
            <div class="col">
                <?= $network['router'] ?>
            </div>
        </div>
    </div>

</div>
<?php endforeach ?>


<?= $this->endSection() ?>