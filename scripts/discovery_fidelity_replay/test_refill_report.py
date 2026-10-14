import unittest
from refill_report import rates, expected, age_bucket

class RefillReportTest(unittest.TestCase):
    def test_missing_exposure_support_is_not_zero_quality(self):
        with self.assertRaises(ValueError):
            expected('1', {('ugc','over30'): 2}, rates([]))

    def test_zero_delivery_agent_and_source_age_sensitivity(self):
        model = rates([
            {'agent_id':'1','lane':'ugc','age':'under7','deliveries':100,'positives':20},
            {'agent_id':'1','lane':'ugc','age':'over30','deliveries':100,'positives':2},
        ])
        self.assertEqual(expected('2', {}, model), (0.0, 0.0))
        source, age = expected('2', {('ugc','over30'): 10}, model)
        self.assertAlmostEqual(source, 1.1)
        self.assertAlmostEqual(age, .2)
        self.assertEqual(age_bucket(7), 'under7')
        self.assertEqual(age_bucket(30), '7to30')
